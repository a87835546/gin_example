#!/bin/sh

# 假如需要查找go脚本main是否运行 当前脚本名称为Status.sh  运行：sh Status.sh 'main'
# 如果需要操作 则第二个参数为：start | restart | stop ｜ build
#项目根目录
ShellPath=$(pwd)
Name=$1

startTask(){
  # shellcheck disable=SC2164
  cd "${ShellPath}"
  nohup ./"${Name}"  >/dev/null 2>&1 &
  echo "starting..."
}
goBuild(){
  # shellcheck disable=SC2164
  cd "${ShellPath}"
  echo "go build begin!"
  go build -o "${Name}"
  echo "go build Success!"
}


# shellcheck disable=SC2039
if [[ $1 == "" ]]
then
  echo "参数错误!如 start.sh main restart"
  exit
fi


#查找进程id
# shellcheck disable=SC2006
# shellcheck disable=SC2009
pid=`ps -ef | grep "$1" | grep -v 'grep' | grep -v "$0" | awk '{print $2}'`
# shellcheck disable=SC2039
if [[ $2 == "start" ]]
then
  if [[ $pid != "" ]]
  then
      echo "$1" "is running"
    else
      startTask
      echo "$1" "Start Success..."
  fi
elif [[ $2 == "stop" ]]
then
  if [[ $pid != "" ]]
   then
       kill "${pid}"
       echo "$1" "Stop Success..."
     else
       echo "$1" "is Stop"
  fi
elif [[ $2 == "restart" ]]
then
   if [[ $pid != "" ]]
        then
            kill "${pid}"
   fi
   startTask
   echo "$1" "Restart Success..."
elif [[ $2 == "build" ]]
then
  goBuild
  if [[ $pid != "" ]]
     then
         kill "${pid}"
  fi
  startTask
else
    if [[ $pid == "" ]]
    then
        echo "$1" "Check Res: Not Running"
        startTask
        echo "$1" "restart Success!"
    else
        echo "$1" "Check Res: Running"
    fi
fi
