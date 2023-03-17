#!/bin/bash

CONTAINER_NAME=cowstudio/be-sample

bash build.sh

# 关闭容器
echo "removing....."
docker stop $CONTAINER_NAME
docker rm $CONTAINER_NAME

# 运行容器
echo "starting....."
docker run -p 3000:8080 -d $CONTAINER_NAME
