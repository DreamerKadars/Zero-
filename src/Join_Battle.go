package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Join_Battle(c *gin.Context) {
	data := Get_cookie(c)

	c.Request.ParseForm()
	uid_str := data["uid"].(string)
	uid, _ := strconv.Atoi(uid_str)

	Boss_id_str := c.Request.Form["Boss_id"][0]
	Boss_id, _ := strconv.Atoi(Boss_id_str)

	err := DB_join_battle(uid, Boss_id)

	if err != nil {
		fmt.Println(err.Error())
		data["result"] = err.Error()
	} else {
		data["result"] = "成功参与！"
	}
	data["nextpage_flag"] = true
	data["nextpage"] = "index"
	c.HTML(http.StatusOK, "verify.html", data)
}

func Exit_verify(c *gin.Context) {
	c.Request.ParseForm()

	uid_str := c.Request.Form["uid"][0]
	uid, _ := strconv.Atoi(uid_str)
	Boss_id_str := c.Request.Form["Boss_id"][0]
	Boss_id, _ := strconv.Atoi(Boss_id_str)

	err := DB_exit_battle(uid, Boss_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	data := gin.H{}
	data["result"] = "退出与" + strconv.Itoa(Boss_id) + "的战斗成功！"
	data["nextpage_flag"] = true
	data["nextpage"] = "index"

	c.HTML(http.StatusOK, "verify.html", data)
}
