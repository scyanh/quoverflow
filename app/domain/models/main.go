package model

import "github.com/gin-gonic/gin"

const (
	StatusSuccess = "success"
	StatusFail    = "fail"
)

type ResponseContext struct {
	Ctx *gin.Context
}

type Response struct {
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
}

type ServiceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (self *ResponseContext) ResponseData(status string, message string, data interface{}) Response {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return response
}