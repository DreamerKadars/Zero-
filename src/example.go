package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("../view/*")

	r.GET("/index", Index)
	r.GET("/someJSON", SomeJSON)

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
