package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	data := Get_cookie(c)

	c.Request.ParseForm()
	if data["login_flag"].(bool) == false {
		//没登陆呢
		c.HTML(http.StatusOK, "index.html", data)
	} else {
		//登陆处理信息
		uid, _ := strconv.Atoi(data["uid"].(string))
		user_battle := DB_get_Battle(uid)
		U_D, _ := DB_get_User_data(uid)
		if U_D.Uid == 0 {
			data["Data_falg"] = false
		} else {
			data["Data_falg"] = true
		}

		data["User_data"] = U_D
		if user_battle == nil {
			data["battle"] = false
		} else {
			data["battle"] = true
			data["user_battle"] = user_battle
		}
		data["User_history"] = DB_get_History(uid)
		data["Boss_Data"], _ = DB_get_Boss_Data()

		c.HTML(http.StatusOK, "index.html", data)
	}

}
