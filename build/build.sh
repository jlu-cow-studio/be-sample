#!/bin/bash

# 替换为国内可用代理
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

# 构建应用
go build -o app main.go

# 安装sidecar
export PATH=$PATH:`go env GOPATH`/bin
go install github.com/jlu-cow-studio/sidecar-service@latest