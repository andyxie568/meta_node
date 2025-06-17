package utils

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ErrorJSON(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{
		Status:  "error",
		Message: message,
	})
}

func SuccessJSON(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status": "success",
		"data":   data,
	})
}
