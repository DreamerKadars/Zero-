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
	c.HTML(http.StatusOK, "verify.html", data)
}
