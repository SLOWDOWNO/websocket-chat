#!/bin/bash

# 端口号
PORT=8080

# 找出占用特定端口的进程
PID=$(lsof -t -i:$PORT)

if [ -z "$PID" ]
then
  echo "没有找到占用端口$PORT的进程."
else
  echo "结束占用端口$PORT的进程，PID为$PID..."
  kill -9 $PID
  echo "进程已结束."
fi