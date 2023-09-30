#!/bin/bash
file=$1

echo "------------$(date +%F' '%T )------------"
echo "------------${file}------------"
# 开发重启脚本


getPid() {
  docmd=$(ps aux | grep ${file} | grep ${file} | grep -v 'grep' | grep -v '\.sh' | awk '{print $2}')
  echo "$docmd"
}

start() {
  pidstr=$(getPid)
    printf "\n"
    printf "正在执行启动...稍候"
    printf "\n"
    nohup ./"${file}" >logs/$(date +%F' '%T ).log 2>&1 &
    pidstr=$(getPid)
    echo "start with pids $pidstr Successful"

}

stop() {

  pidstr=$(getPid)
  if [ ! -n "$pidstr" ]; then
    echo "Not Executed!"
    return
  fi

  echo "kill $pidstr done"
  kill $pidstr

}

restart() {
  stop
  start
}

case "$2" in
start)
  start
  ;;
stop)
  stop
  ;;
restart)
  restart
  ;;
getpid)
  getPid
  ;;
esac