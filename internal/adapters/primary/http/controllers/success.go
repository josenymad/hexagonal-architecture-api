package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Data   map[string]interface{} `json:"data"`
	Status string                 `json:"status"`
}

type SuccessController struct {
}

func NewSuccessController() SuccessController {

	return SuccessController{}
}

func getRequestData(c *gin.Context) ResponseData {
	requestData := make(map[string]interface{})
	c.BindJSON(&requestData)
	queryParams := c.Request.URL.Query()

	for key, values := range queryParams {
		requestData[key] = values[0]
	}

	return ResponseData{
		Data:   requestData,
		Status: requestData["status"].(string),
	}
}

func (sc *SuccessController) SuccessHandler(c *gin.Context) {
	responseData := getRequestData(c)
	responseStatus, err := strconv.Atoi(responseData.Status)
	if err != nil {
		fmt.Println("Could not convert status from string to int", err)
		return
	}
	c.JSON(responseStatus, responseData.Data)
}
