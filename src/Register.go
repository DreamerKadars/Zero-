package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "Main website"})
}
