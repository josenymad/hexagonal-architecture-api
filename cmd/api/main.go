package main

import (
	"github.com/josenymad/hexagonal-go-api/internal/adapters/primary/http"
	"github.com/josenymad/hexagonal-go-api/internal/core/services"
)

func main() {
	service := services.NewService()
	http := http.NewHttpServer("localhost", "8000")
	http.Run(service)
}
