# 基础镜像
FROM golang:1.17-alpine

# 设置工作目录
WORKDIR /app

# 将宿主机中的当前目录的所有文件拷贝到镜像中的 /app 目录下
COPY . .

# 替换为国内可用代理
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct

# 构建应用
RUN go build -o app main.go

# 暴露服务端口
EXPOSE 8080

# 执行命令
CMD ["/app/app"]
