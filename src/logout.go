package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	fmt.Println("注销账号")
	c.SetCookie("uid", "", -1, "", "127.0.0.1", false, true)
	c.Redirect(http.StatusMovedPermanently, "login")
}
