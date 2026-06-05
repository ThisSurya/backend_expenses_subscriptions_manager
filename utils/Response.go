package utils

import "github.com/gin-gonic/gin"

type APIResponse struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func SuccessResponse(c *gin.Context, message string, data interface{}, status int) {
	c.JSON(status, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, message string, errors interface{}, status int) {
	c.JSON(status, APIResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}
