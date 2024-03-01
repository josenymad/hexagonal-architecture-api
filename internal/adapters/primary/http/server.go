package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/josenymad/hexagonal-go-api/internal/adapters/primary/http/controllers"
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

func (server *HttpServer) Run() {
	successController := controllers.NewSuccessController()
	router := gin.New()
	router.POST("/success", successController.SuccessHandler)
	router.Run(fmt.Sprintf("%s:%s", server.address, server.port))
}
