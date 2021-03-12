package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {

	fmt.Print("服务器启动！")
	go Run_by_Time(10)

	r := gin.Default()
	r.LoadHTMLGlob("../view/*")
	r.Static("./static", "../static")

	r.GET("/", Index)
	r.GET("/index", Index)                   //主页
	r.GET("/register", Register)             //注册界面
	r.GET("/login", Login)                   //登陆界面
	r.GET("/logout", Logout)                 //注销
	r.GET("/admpage", AdmPage)               //管理员
	r.GET("/MyPage", MyPage)                 //用户主页
	r.GET("/Boss_data_page", Boss_data_page) //Boss信息一览
	r.GET("/BattlePage", BattlePage)         //战斗界面
	r.GET("/someJSON", SomeJSON)

	r.POST("/register_verify", Register_verify)                 //注册验证
	r.POST("/login_verify", Login_verify)                       //登陆验证
	r.POST("/Init_User_data", Init_User_data)                   //更新用户数据
	r.POST("/Boss_data_add", Boss_data_add)                     //插入boss信息
	r.POST("/Join_Battle", Join_Battle)                         //参与战斗信息
	r.POST("/Hit_verify", Hit_verify)                           //Hit验证信息
	r.POST("/Exit_verify", Exit_verify)                         //退出验证
	r.POST("/user_num_add", user_num_add)                       //添加一定数量的用户
	r.POST("/Quick_Join_And_Hit_Html", Quick_Join_And_Hit_Html) //快速出击验证
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Use(TlsHandler())
	r.RunTLS(":443", "../static/5305314_www.loveranran.xyz.pem", "../static/5305314_www.loveranran.xyz.key")

}

func TlsHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
