package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AdmPage(c *gin.Context) {
	data := Get_cookie(c)
	uid, _ := strconv.Atoi(data["uid"].(string))
	err := DB_is_adm(uid)
	if err == nil {
		data["Adm"] = true
	} else {
		data["Adm"] = false
	}

	var DB_data [6]string
	for i := 0; i < 6; i++ {
		DB_data[i] = DB_name[i] + "的数量是：" + strconv.Itoa(DB_num(DB_name[i]))
	}

	data["DB_information"] = DB_data

	c.HTML(http.StatusOK, "AdmPage.html", data)
}
