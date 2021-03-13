dtm=`date +"%Y年%m月%d日%H时%M分%S秒"`
echo $dtm
start_time=$(date +%s)

test_file=./uid_pwd.txt

rowCnt=`cat $test_file | wc -l `

# 读取文件中的内容到数组中
array=($(cat $test_file))

num=10
for((v=1;v<=$num;v++))
do
    curl -s -w '\n' -d ${array[v]} -k https://www.loveranran.xyz/Quick_Join_And_Hit_Shell >> result.txt &
done

wait #等待子任务全部完成

end_time=$(date +%s)
cost_time=$[ $end_time-$start_time ]
echo "成功完成 $num 个请求"
echo "总共用时 $cost_time s"
num3=`echo "scale=2; $cost_time/$num" | bc`
echo "平均用时 $num3 s"
