package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func user_num_add(c *gin.Context) {
	var data = gin.H{}
	c.Request.ParseForm()
	user_num, _ := strconv.Atoi(c.Request.Form["user_num"][0])

	succ_num := DB_add_user(user_num) //返回成功数量

	data["result"] = "成功插入了" + strconv.Itoa(succ_num) + "个用户！"
	data["nextpage_flag"] = true
	data["nextpage"] = "admpage"
	c.HTML(http.StatusOK, "verify.html", data)
}
