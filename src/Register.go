package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "Main website"})
}
func Register_verify(c *gin.Context) {
	fmt.Println("接受用户注册请求")
	c.Request.ParseForm()
	fmt.Println("收到的数据", c.Request.Form)
	//c.Request.Form[];
	c.HTML(http.StatusOK, "register_verify.html", c.Request.Form)
}
