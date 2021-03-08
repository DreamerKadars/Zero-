package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Get_cookie(c *gin.Context) map[string]interface{} {
	var data map[string]interface{} = make(map[string]interface{})
	uid, err := c.Cookie("uid")

	fmt.Println("获取cookie ", c.Request.Header["Cookie"])

	if err != nil {
		data["login_flag"] = false
		return data
	}
	data["uid"] = uid
	data["login_flag"] = 1
	return data
}
