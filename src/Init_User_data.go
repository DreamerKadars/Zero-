package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Init_User_data(c *gin.Context) {

	data := Get_cookie(c)

	var U_D User_data
	c.Request.ParseForm()
	fmt.Println(c.Request.Form)
	U_D.Name = c.Request.Form["name"][0]
	U_D.Mola = 0
	rand.Seed(time.Now().Unix())
	U_D.Atk = 30 + rand.Intn(10)
	U_D.Uid, _ = strconv.Atoi(data["uid"].(string))
	U_D.Buff1 = 0
	U_D.Buff2 = 0
	U_D.Buff3 = 0

	data["User_data"] = U_D
	err := DB_insert_User_data(U_D)
	if err != nil {
		fmt.Println(err)
		return
	}
	data["result"] = "成功初始化信息！"
	c.HTML(http.StatusOK, "verify.html", data)
}
