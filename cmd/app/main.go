package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/itoqsky/money-tracker-backend/internal/service"
	"github.com/itoqsky/money-tracker-backend/internal/storage"
	"github.com/itoqsky/money-tracker-backend/internal/transport/rest"
	"github.com/itoqsky/money-tracker-backend/internal/transport/rest/handler"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error occured while initializing configs: %s", err.Error())
	}

	db, err := storage.NewPostgresDB(storage.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("error occured while connecting to db: %s", err.Error())
	}

	storages := storage.NewStorage(db)
	services := service.NewService(storages)
	handlers := handler.NewHandler(services)

	srv := new(rest.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("\nMoney-tracker app shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on closing database: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
