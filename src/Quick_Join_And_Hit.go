package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	B_D, _ := DB_Find_Min_Hp_Boss(1)
	if B_D == nil {
		fmt.Println("寻找Boss失败！！")
		return
	}

	Boss_id := B_D[0].Boss_id
	Mola := B_D[0].Mola
	Atk := U_D.Atk

	_, result := Hit_Boss(uid, Boss_id, Atk, Mola)

	data["result"] = result

	data["nextpage_flag"] = true
	data["nextpage"] = "index"

	c.HTML(http.StatusOK, "verify.html", data)
}

const Quick_Boss_num int = 5

func Quick_Join_And_Hit_Shell(c *gin.Context) {
	data := gin.H{}
	var result string
	//拿到post的uid和pwd，验证
	c.Request.ParseForm()
	var pwd = c.Request.Form["pwd"][0]
	uid_str := c.Request.Form["uid"][0]
	uid, _ := strconv.Atoi(uid_str)
	err := DB_found(uid, pwd)
	if err != nil {
		result = "失败，用户名或者密码错误"
	} else {
		//找一下用户信息
		U_D, err := DB_get_User_data(uid)
		if err != nil {
			fmt.Println("寻找User失败！！")
			fmt.Println(err.Error())
			return
		}
		//选择血量最少的几个boss
		D_B, err := DB_Find_Min_Hp_Boss(Quick_Boss_num)
		if err != nil {
			result = "失败，找不到目标"
			return
		}
		//随机选择一个Hit
		rand.Seed(time.Now().Unix())
		var boss_flag int = rand.Intn(Quick_Boss_num)
		_, result = Hit_Boss(uid, D_B[boss_flag].Boss_id, U_D.Atk, D_B[boss_flag].Mola)
		//等待结果
	}

	data["result"] = result
	c.HTML(http.StatusOK, "Simple_Replay.html", data)
}
