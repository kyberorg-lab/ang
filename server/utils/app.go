package utils

import "github.com/gin-gonic/gin"

type HttpUtils struct {
	Context *gin.Context
}

func (u *HttpUtils) SetJsonHeader() {
	u.Context.Header("Content-Type", "application/json")

}

func (u *HttpUtils) SendError(code int, message string) {
	u.Context.JSON(code, gin.H{"error": message})
}
