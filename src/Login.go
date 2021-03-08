package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func Login_verify(c *gin.Context) {
	fmt.Println("接受用户登陆请求")
	c.Request.ParseForm()
	fmt.Println("收到的数据", c.Request.Form)
	var pwd = c.Request.Form["pwd"][0]
	uid, _ := strconv.Atoi(c.Request.Form["uid"][0])

	err := DB_found(uid, pwd)
	//c.Request.Form[];
	var result string
	if err == nil {
		result = "尊敬的用户" + c.Request.Form["uid"][0] + "，恭喜您登陆成功"
	} else {
		result = err.Error()
	}
	data := map[string]interface{}{
		"result": result,
	}
	c.HTML(http.StatusOK, "login_verify.html", data)
}
