package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	api "go-api"
	cfg "go-api/config"
	"go-api/pkg/handler"
	"go-api/pkg/repository"
	"go-api/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	config, err := initConfig()
	if err != nil {
		logrus.Fatalf("Error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(config)
	if err != nil {
		logrus.Fatalf("Failed to initialize DB: %s\n", err.Error())
	}

	handlers := handler.NewHandler(service.NewService(repository.NewRepository(db)))

	srv := new(api.Server)
	if err := srv.Run(config.AppPort, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while running http server: %s\n", err.Error())
	}
}

func initConfig() (*cfg.Config, error) {
	return new(cfg.Config).Init("config/config.local.yml")
}
