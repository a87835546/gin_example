#! /bin/sh
cd main/
ls
printf "输入你想操作的内容 \n 1. 构建linux的项目 \n 2.运行项目 "
read num
case $num in
  1)
    # 构建linux的运行包
    CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build main.go
  ;;
 2)
    # 运行项目
    go run main.go
  ;;
  *)
    echo "其他的输出，不知道你的想法"
esac
