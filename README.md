# Gin 最小服务

一个最小可用的 Gin Web 服务示例。

## 功能

- `/` - 根路径，返回欢迎消息
- `/health` - 健康检查端点

## 本地运行

```bash
# 安装依赖
go mod download

# 运行服务
go run main.go
```

服务将在 `http://localhost:8080` 启动

## Docker 构建和运行

```bash
# 构建镜像
docker build -t gin-server .

# 运行容器
docker run -p 8080:8080 gin-server
```

## 测试

```bash
# 测试根路径
curl http://localhost:8080/

# 测试健康检查
curl http://localhost:8080/health
```
