package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Data   map[string]interface{} `json:"data"`
	Status string                 `json:"status"`
}

type TestController struct {
}

func NewTestController() TestController {

	return TestController{}
}

func getRequestData(c *gin.Context) ResponseData {
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

	return ResponseData{
		Data:   requestData,
		Status: status,
	}
}

func (sc *TestController) TestHandler(c *gin.Context) {
	responseData := getRequestData(c)
	responseStatus, err := strconv.Atoi(responseData.Status)
	if err != nil {
		log.Println("Error converting status from string to int", err)
		return
	}
	c.JSON(responseStatus, responseData.Data)
}