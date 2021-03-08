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
	uid_str := c.Request.Form["uid"][0]
	uid, _ := strconv.Atoi(uid_str)
	err := DB_found(uid, pwd)
	//c.Request.Form[];
	var result string
	if err == nil {
		//设置cookie
		fmt.Println(uid_str + "登陆成功，设置cookie！")
		c.SetCookie("uid", uid_str, 0, "/", "127.0.0.1", false, true) //时间为0,代表是会话cookie
		result = "尊敬的用户" + uid_str + "，恭喜您登陆成功"
	} else {
		result = err.Error()
	}
	data := Get_cookie(c)
	data["result"] = result
	c.HTML(http.StatusOK, "login_verify.html", data)
}
