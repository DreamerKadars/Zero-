package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BattlePage(c *gin.Context) {
	data := Get_cookie(c)
	c.Request.ParseForm()
	//用户信息
	uid, _ := strconv.Atoi(data["uid"].(string))
	U_D, _ := DB_get_User_data(uid)
	data["User_data"] = U_D
	data["User_data_str"] = Get_User_data_str(U_D)

	if c.Request.Form["Boss_id"] == nil {
		return
	}
	//Boss信息
	var t = c.Request.Form["Boss_id"][0]
	Boss_id, _ := strconv.Atoi(t)

	B_D, _ := DB_get_Boss_Data_one(Boss_id)
	if B_D == nil {
		return
	}
	data["Boss_data"] = B_D[0]
	data["Boss_data_str"] = Get_Boss_data_str(B_D[0])

	var U_D_Complete = DB_Compete(uid, Boss_id)
	var U_D_Complete_str []string = make([]string, len(U_D_Complete))
	for key, value := range U_D_Complete {
		U_D_Complete_str[key] = strconv.Itoa(value.Uid) + "  :  " + value.Name
	}
	fmt.Println(U_D_Complete_str)
	data["U_D_Complete_str"] = U_D_Complete_str

	data["Boss_id"] = Boss_id
	data["Atk"] = U_D.Atk
	c.HTML(http.StatusOK, "BattlePage.html", data)
}

func Hit_verify(c *gin.Context) {
	data := Get_cookie(c)
	c.Request.ParseForm()
	uid, _ := strconv.Atoi(data["uid"].(string))
	Boss_id, _ := strconv.Atoi(c.Request.Form["Boss_id"][0])
	Atk, _ := strconv.Atoi(c.Request.Form["Atk"][0])
	var Re chan bool = make(chan bool)
	var T Hit = Hit{uid: uid, boss_id: Boss_id, atk: Atk, Re_chan: Re}
	Hit_ch <- T
	var flag = <-Re
	var result string

	if flag == true {
		result = "出击成功"
	} else {
		result = "出击失败"
	}
	data["result"] = result
	c.HTML(http.StatusOK, "verify.html", data)
}
