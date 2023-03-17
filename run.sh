#!/bin/bash

bash build.sh

# 运行容器
docker run -p 3000:8080 -d cowstudio/be-sample
