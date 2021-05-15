package main

import (
	api "go-api"
	cfg "go-api/config"
	"go-api/pkg/handler"
	"go-api/pkg/repository"
	"go-api/pkg/service"
	"log"
)

func main() {
	config, err := initConfig()
	if err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}

	handlers := handler.NewHandler(service.NewService(repository.NewRepository()))

	srv := new(api.Server)
	if err := srv.Run(config.Port, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s\n", err.Error())
	}
}

func initConfig() (*cfg.Config, error) {
	return new(cfg.Config).Init("config/config.yml")
}
