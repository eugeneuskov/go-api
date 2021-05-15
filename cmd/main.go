package main

import (
	_ "github.com/lib/pq"
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

	log.Printf("postgres config: %v\n", config.Db.Postgres.Host)

	db, err := repository.NewPostgresDB(config)
	if err != nil {
		log.Fatalf("Failed to initialize DB: %s\n", err.Error())
	}

	handlers := handler.NewHandler(service.NewService(repository.NewRepository(db)))

	srv := new(api.Server)
	if err := srv.Run(config.AppPort, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s\n", err.Error())
	}
}

func initConfig() (*cfg.Config, error) {
	return new(cfg.Config).Init("config/config.local.yml")
}
