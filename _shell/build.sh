#!/bin/bash
# 加载变量
source "./_shell/init.sh"
#############

echo " =========== go build  =========== "

go mod tidy &&
  go build -o ${buildName}
echo " server 端编译 完成"

echo " =========== 开始进行 文件整理 =========== "

echo "移动 go build 文件"
mv ${buildName} ${outPutPath} &&
  exit
