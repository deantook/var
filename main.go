package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	//从系统变量中获取名为 hello 的变量
	hello := os.Getenv("hello")
	fmt.Println("hello", hello)
	fmt.Println("hello", hello)
	fmt.Println("hello", hello)
	fmt.Println("hello", hello)
	fmt.Println("hello", hello)
	fmt.Println("hello", hello)
	fmt.Println("hello", hello)
	fmt.Println("hello", hello)
	fmt.Println("hello", hello)
	fmt.Println("hello", hello)

	// 打印日志

	log.Println("hello", hello)
	log.Println("hello", hello)
	log.Println("hello", hello)
	log.Println("hello", hello)
	log.Println("hello", hello)
	log.Println("hello", hello)

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

	//从系统变量中获取名为 hello 的变量
	hello1 := os.Getenv("hello")
	fmt.Println("hello", hello1)
	fmt.Println("hello", hello1)
	fmt.Println("hello", hello1)
	fmt.Println("hello", hello1)
	fmt.Println("hello", hello1)
	fmt.Println("hello", hello1)
	fmt.Println("hello", hello1)
	fmt.Println("hello", hello1)
	fmt.Println("hello", hello1)
	fmt.Println("hello", hello1)

	// 打印日志1

	log.Println("hello", hello1)
	log.Println("hello", hello1)
	log.Println("hello", hello1)
	log.Println("hello", hello1)
	log.Println("hello", hello1)
	log.Println("hello", hello1)

}
