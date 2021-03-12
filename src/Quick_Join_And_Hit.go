package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Quick_Join_And_Hit_Html(c *gin.Context) {
	//拿到cookie的uid
	data := Get_cookie(c)
	uid, _ := strconv.Atoi(data["uid"].(string))
	U_D, err := DB_get_User_data(uid)
	if err != nil {
		fmt.Println("寻找User失败！！")
		fmt.Println(err.Error())
		return
	}
	//选择一个血量最少的boss
	B_D, _ := DB_Find_Min_Hp_Boss()
	if B_D == nil {
		fmt.Println("寻找Boss失败！！")
		return
	}

	Boss_id := B_D[0].Boss_id
	Mola := B_D[0].Mola
	Atk := U_D.Atk

	var Re chan int = make(chan int)

	var T Hit = Hit{uid: uid, boss_id: Boss_id, atk: Atk, Re_chan: Re}

	//进攻
	DB_join_battle(uid, Boss_id)
	Hit_ch <- T

	var flag = <-Re

	var result string

	if flag != 0 {
		result = "出击成功，对Boss：" + strconv.Itoa(Boss_id) + "造成了" + strconv.Itoa(Atk) + "点伤害，当前Boss还剩余" + strconv.Itoa(flag) + "点血量"
	} else if flag == -1 {
		result = "出击失败"
	} else if flag == 0 {
		result = "成功击杀Boss，获得赏金" + strconv.Itoa(Mola)
	}

	data["result"] = result

	data["nextpage_flag"] = true
	data["nextpage"] = "index"

	c.HTML(http.StatusOK, "verify.html", data)
}

func Quick_Join_And_Hit_Shell(c *gin.Context) {
	//拿到post的uid和pwd，验证
	//选择血量最少的几个boss
	//随机选择一个Hit
	//等待结果
}
