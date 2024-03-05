package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/josenymad/hexagonal-go-api/internal/adapters/primary/http/controllers"
	"github.com/josenymad/hexagonal-go-api/internal/core/services"
)

type HttpServer struct {
	address string
	port    string
}

func NewHttpServer(address, port string) HttpServer {
	return HttpServer{
		address: address,
		port:    port,
	}
}

func (server *HttpServer) Run(service services.Service) {
	testController := controllers.NewController(service)
	router := gin.New()
	router.POST("/test", testController.TestHandler)
	router.Run(fmt.Sprintf("%s:%s", server.address, server.port))
}
