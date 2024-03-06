package main

import (
	"log"

	"github.com/josenymad/hexagonal-go-api/internal/adapters/primary/http"
	"github.com/josenymad/hexagonal-go-api/internal/adapters/secondary/database/mongo"
	"github.com/josenymad/hexagonal-go-api/internal/core/services"
)

func main() {
	mongoDB, err := mongo.NewMongoDB("hexagonalDB", "hexagonal")
	if err != nil {
		log.Fatalln("Failed to connect to mongoDB", err)
	}
	service := services.NewService(&mongoDB)
	http := http.NewHttpServer("localhost", "8000")
	http.Run(service)
}
