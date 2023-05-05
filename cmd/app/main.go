package main

import (
	"log"

	"github.com/itoqsky/money-tracker-backend/internal/service"
	"github.com/itoqsky/money-tracker-backend/internal/storage"
	"github.com/itoqsky/money-tracker-backend/internal/transport/rest"
	"github.com/itoqsky/money-tracker-backend/internal/transport/rest/handler"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occured while initializing configs: %s", err.Error())
	}

	storages := storage.NewStorage()
	services := service.NewService(storages)
	handlers := handler.NewHandler(services)

	srv := new(rest.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
