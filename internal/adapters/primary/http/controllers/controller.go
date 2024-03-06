package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/josenymad/hexagonal-go-api/internal/core/services"
)

type Response struct {
	Data   map[string]interface{} `json:"data"`
	Status string                 `json:"status"`
}

type Controller struct {
	service services.Service
}

func NewController(service services.Service) Controller {

	return Controller{
		service: service,
	}
}

func getRequestData(c *gin.Context) Response {
	requestData := make(map[string]interface{})

	if err := c.BindJSON(&requestData); err != nil {
		log.Println("Error binding request data JSON", err)
	}

	queryParams := c.Request.URL.Query()

	for key, values := range queryParams {
		requestData[key] = values[0]
	}

	var status string

	if requestData["status"] != nil {
		status = requestData["status"].(string)
	} else {
		status = "200"
	}

	return Response{
		Data:   requestData,
		Status: status,
	}
}

func (controller *Controller) Handler(c *gin.Context) {
	response := getRequestData(c)
	responseStatus, err := strconv.Atoi(response.Status)
	if err != nil {
		log.Println("Error converting status from string to int", err)
		return
	}

	controller.service.SendData(response.Data)

	c.JSON(responseStatus, response.Data)
}
