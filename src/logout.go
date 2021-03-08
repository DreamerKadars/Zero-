package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.SetCookie("uid", "", -1, "", "", false, true)
	c.Redirect(http.StatusMovedPermanently, "login")
}
