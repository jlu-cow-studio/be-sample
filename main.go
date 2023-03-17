package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建gin引擎
	router := gin.Default()

	// 定义路由及处理函数
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	// 启动HTTP服务，默认监听地址为:8080
	router.Run("0.0.0.0:3001")
}
