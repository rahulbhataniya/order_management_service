package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendSuccessResponse sends a success response
func SendSuccessResponse(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

// SendErrorResponse sends an error response
func SendErrorResponse(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
}
