package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
	router := gin.New()
	//router.POST("/success")
	router.Run(fmt.Sprintf("%s:%s", server.address, server.port))
}
