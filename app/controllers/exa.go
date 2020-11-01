package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 接口连接测试
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
