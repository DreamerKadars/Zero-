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
	var User_data_str [7]string
	User_data_str[0] = "Uid  :" + strconv.Itoa(U_D.Uid)
	User_data_str[1] = "Name :" + U_D.Name
	User_data_str[2] = "ATK  :" + strconv.Itoa(U_D.Atk)
	User_data_str[3] = "Mola :" + strconv.Itoa(U_D.Mola)
	User_data_str[4] = "状态1 :" + strconv.Itoa(U_D.Buff1)
	User_data_str[5] = "状态2 :" + strconv.Itoa(U_D.Buff2)
	User_data_str[6] = "状态3 :" + strconv.Itoa(U_D.Buff3)
	data["User_data_str"] = User_data_str
	c.HTML(http.StatusOK, "MyPage.html", data)
}
