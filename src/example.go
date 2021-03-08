package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("服务器启动！")
	r := gin.Default()
	r.LoadHTMLGlob("../view/*")
	r.Static("./static", "../static")
	r.GET("/index", Index)
	r.GET("/register", Register) //注册界面
	r.GET("/login", Login)       //登陆界面
	r.GET("/logout", Logout)     //注销
	r.GET("/someJSON", SomeJSON)
	r.POST("/register_verify", Register_verify) //注册验证
	r.POST("/login_verify", Login_verify)       //登陆验证
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")

}
