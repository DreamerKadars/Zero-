package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", Get_cookie(c))
}
func Register_verify(c *gin.Context) {
	fmt.Println("接受用户注册请求")
	c.Request.ParseForm()
	fmt.Println("收到的数据", c.Request.Form)
	var pwd = c.Request.Form["pwd"][0]
	uid, _ := strconv.Atoi(c.Request.Form["uid"][0])

	err := DB_register_point(uid, pwd, DB)
	//c.Request.Form[];
	var result string
	if err == nil {
		result = "尊敬的用户" + c.Request.Form["uid"][0] + "，恭喜您注册成功"
	} else {
		result = err.Error()
	}
	data := Get_cookie(c)
	data["result"] = result

	data["nextpage_flag"] = true
	data["nextpage"] = "MyPage"

	c.HTML(http.StatusOK, "verify.html", data)
}
