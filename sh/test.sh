#!/bin/bash
Njob=100 #任务总数
Nproc=5 #最大并发进程数
function PushQue { #将pid追加到队列中
    Que="$Que $1"
    Nrun=$(($Nrun+1))
    #echo $Nrun
}
function GenQue { #更新队列信息，先清空队列消息，然后检索生成新的消息队列
    OldQue=$Que
    Que="";Nrun=0
    for PID in $OldQue ; do
        ps -p $PID > /dev/null
        if [[ $? -eq 0 ]]; then
            PushQue $PID
        fi
    done
}
function ChkQue { #检查队列消息，如果有已经结束了的进程的PID，那么更新队列信息
    OldQue=$Que
    for PID in $OldQue; do
        ps -p $PID > /dev/null
        if [[ ! $? -eq 0 ]]; then
            GenQue; break
        fi
    done
}
for ((i=0; i<$Njob; i++)); do
    ./Quick_Hit.sh & #循环内容放到后台执行
    PID=$!
    PushQue $PID
    while [[ $Nrun -ge $Nproc ]]; do #如果Nrun大于Nproc，就一直Chkque
        ChkQue
        sleep 0.1
    done
done
wait #等待子任务全部完成

echo -e "time-consuming: $SECONDS seconds" #显示脚本执行耗时