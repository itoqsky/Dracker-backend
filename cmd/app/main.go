package main

import (
	"log"

	"github.com/itoqsky/money-tracker-backend/internal/transport/rest"
	"github.com/itoqsky/money-tracker-backend/internal/transport/rest/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(rest.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
