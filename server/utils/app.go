package utils

import "github.com/gin-gonic/gin"

func SetJsonHeader(context *gin.Context) {
	context.Header("Content-Type", "application/json")

}

func SendError(context *gin.Context, code int, message string) {
	context.JSON(code, gin.H{"error": message})
}
