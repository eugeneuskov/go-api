package main

import (
	api "go-api"
	"go-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(api.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s\n", err.Error())
	}

}
