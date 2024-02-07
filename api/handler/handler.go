package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test/api/models"
	"test/service"
)

type Handler struct {
	services service.IServiceManager
}

func New(services service.IServiceManager) Handler {
	return Handler{
		services: services,
	}
}

func handleResponse(c *gin.Context, msg string, statusCode int, data interface{}) {
	resp := models.Response{}

	switch code := statusCode; {
	case code < 400:
		resp.Description = "success"
	case code < 500:
		resp.Description = "bad request"
		fmt.Println("BAD REQUEST: "+msg, "reason: ", data)
	default:
		resp.Description = "internal server error"
		fmt.Println("INTERNAL SERVER ERROR: "+msg, "reason: ", data)
	}

	resp.StatusCode = statusCode
	resp.Data = data

	c.JSON(resp.StatusCode, resp)
}