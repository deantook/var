package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	// 从环境变量获取 Redis URL
	redisURL := os.Getenv("REDIS_URL")
	if redisURL != "" {
		log.Printf("检测到 REDIS_URL: %s", redisURL)
		
		// 连接 Redis 并设置键值对 hello:worldddd
		client, err := connectRedis(redisURL)
		if err != nil {
			log.Printf("连接 Redis 失败: %v", err)
		} else {
			ctx := context.Background()
			// 设置键值对 hello:worldddd
			err = client.Set(ctx, "hello", "worldddd", 0).Err()
			if err != nil {
				log.Printf("设置 Redis 键值失败: %v", err)
			} else {
				log.Println("成功设置 Redis 键值: hello -> worldddd")
			}
			client.Close()
		}
	} else {
		log.Println("未设置 REDIS_URL 环境变量，跳过 Redis 连接")
	}

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
	log.Println("Gin 服务器启动在 :8080")
	r.Run(":8080")
}

// connectRedis 根据 Redis URL 连接 Redis
func connectRedis(redisURL string) (*redis.Client, error) {
	// 解析 Redis URL
	// 格式: redis://[username:password@]host:port[/db]
	parsedURL, err := url.Parse(redisURL)
	if err != nil {
		return nil, fmt.Errorf("解析 Redis URL 失败: %w", err)
	}

	// 提取用户名和密码
	var username, password string
	if parsedURL.User != nil {
		username = parsedURL.User.Username()
		password, _ = parsedURL.User.Password()
	}

	// 提取主机和端口
	host := parsedURL.Hostname()
	port := parsedURL.Port()
	if port == "" {
		port = "6379" // 默认端口
	}
	addr := fmt.Sprintf("%s:%s", host, port)

	// 提取数据库编号
	db := 0
	if parsedURL.Path != "" {
		dbStr := strings.TrimPrefix(parsedURL.Path, "/")
		if dbStr != "" {
			db, _ = strconv.Atoi(dbStr)
		}
	}

	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
		DB:       db,
	})

	// 测试连接
	ctx := context.Background()
	_, err = client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("Redis 连接测试失败: %w", err)
	}

	return client, nil
}
