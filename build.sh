#!/bin/bash

# 构建 Docker 镜像
docker build -t nerdctl-builder .

# 从容器中复制打包好的文件
docker create --name temp-container nerdctl-builder
docker cp temp-container:/dist/nerdctld.tar.gz .
docker rm temp-container

echo "构建完成！生成了以下文件："
echo "- nerdctld.tar.gz (服务端)"
echo "- nerdctlc.tar.gz (客户端)" 