

# 判断是否启动
# STATUS=ps -aux|grep "./SimpleBlockchain"|awk '$11!="grep"{print $2}'



#!/bin/bash
# . /etc/init.d/functions

devsNum=`ps -eo comm,pid|awk '/SimpleBlock/'|wc -l`
simPID=`ps -aux|grep "./SimpleBlockchain"|awk '$11!="grep"{print $2}'`


start()
{
    echo "start"
    if [ $devsNum == 0 ]; then
        echo "start SimpleBlock"
        ./SimpleBlockchain.exe &
    fi
}

stop()
{
    echo "stop"
      echo "start"
    if [ $devsNum != 0 ]; then
        echo "stop SimpleBlock"
        kill -9 $simPID
    fi
}

status()
{
    echo $devsNum
    echo $simPID
    echo "status"
}

case $1 in 
 start)
      start
      ;;
 stop)
      stop
      ;;
 status)
     status
     ;;
 *)
   echo "Usage $0 {start|stop|status}"
   exit 0
esac