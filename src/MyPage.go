package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MyPage(c *gin.Context) {
	data := Get_cookie(c)
	uid, _ := strconv.Atoi(data["uid"].(string))
	U_D, err := DB_get_User_data(uid)
	if err != nil {
		fmt.Println("第一次使用 ", err)
	}
	fmt.Println("默认用户信息： ", U_D)
	data["User_data"] = U_D

	if U_D.Uid == 0 {
		data["Data_falg"] = false
	} else {
		data["Data_falg"] = true
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
