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
	r.GET("/register", Register)             //注册界面
	r.GET("/login", Login)                   //登陆界面
	r.GET("/logout", Logout)                 //注销
	r.GET("/admpage", AdmPage)               //管理员
	r.GET("/MyPage", MyPage)                 //用户主页
	r.GET("/Boss_data_page", Boss_data_page) //Boss信息一览
	r.GET("/someJSON", SomeJSON)
	r.POST("/register_verify", Register_verify) //注册验证
	r.POST("/login_verify", Login_verify)       //登陆验证
	r.POST("/Init_User_data", Init_User_data)   //更新用户数据
	r.POST("/Boss_data_add", Boss_data_add)     //插入boss信息
	r.POST("/Join_Battle", Join_Battle)         //参与战斗信息
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")

}
