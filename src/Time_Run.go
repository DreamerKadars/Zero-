package main

import "time"

var Clear_Boss_Hp0_time int

//此文件暂时用不上
func Clear_Boss_Hp0(Clear_Boss_Hp0_time int) {
	for {
		time.Sleep(time.Duration(Clear_Boss_Hp0_time) * time.Second)

	}
}
