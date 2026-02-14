# 使用官方 Go 镜像作为构建阶段
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go mod 文件
COPY go.mod go.sum* ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
# CGO_ENABLED=0 禁用 CGO，生成静态链接的二进制文件
# GOOS=linux 指定目标操作系统
# -a 强制重新构建所有包
# -installsuffix cgo 使用 cgo 后缀安装
# -o server 指定输出文件名
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# 使用轻量级 alpine 镜像作为运行阶段
FROM alpine:latest

# 安装 ca-certificates 用于 HTTPS 请求
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制二进制文件
COPY --from=builder /app/server .

# 暴露 8080 端口
EXPOSE 8080

# 运行应用
CMD ["./server"]
