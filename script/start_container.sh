#!/bin/bash
# : bash start_container.sh


# 容器ID或名称
CONTAINER="container-websocket-chat"
CONTAINER_ID="9e3105c0d436"

# 检查容器是否已经在运行
if [ "$(docker ps -q -f name=$CONTAINER)" ]; then
    echo "容器已经在运行."
else
    if [ "$(docker ps -aq -f status=exited -f name=$CONTAINER)" ]; then
        # 如果容器已经存在但是没有运行，那么启动容器
        echo "启动容器..."
        docker start $CONTAINER
    else
        # 如果容器不存在，那么打印错误消息
        echo "容器不存在."
    fi
fi