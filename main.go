package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建 gin 路由引擎
	r := gin.Default()

	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 根路径
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// 启动服务器，监听 8080 端口
	r.Run(":8080")
}
