package main

import (
	"fmt"
	"time"
)

var Clear_Boss_Hp0_time int

const Boss_Live_num_min int = 30

//如果存活的Boss少于Boss_Live_num_min,则增加Boss_Live_num_min个
func Add_Boss(add_num int) {
	Boss_num := DB_get_Live_Boss_num()
	if Boss_num < Boss_Live_num_min {
		DB_insert_Boss_data(Creat_Boss_rand(Boss_Live_num_min))
	}
}

//定时函数
func Run_by_Time(Run_time int) {
	fmt.Println("开启周期为", Run_time, "s的定时任务")
	for {
		time.Sleep(time.Duration(Run_time) * time.Second)
		Add_Boss(Boss_Live_num_min)
	}
}
