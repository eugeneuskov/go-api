package main

import (
	api "go-api"
	"go-api/pkg/handler"
	"go-api/pkg/repository"
	"go-api/pkg/service"
	"log"
)

func main() {
	handlers := handler.NewHandler(service.NewService(repository.NewRepository()))

	srv := new(api.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s\n", err.Error())
	}

}
