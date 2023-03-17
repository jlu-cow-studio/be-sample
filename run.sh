#!/bin/bash

bash build.sh

# 运行容器
docker run -p 3000:3000 -d cowstudio/be-sample
