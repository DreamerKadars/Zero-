package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Boss_data_add(c *gin.Context) {
	c.Request.ParseForm()
	boss_num, _ := strconv.Atoi(c.Request.Form["boss_num"][0])
	var B_D []Boss_data = make([]Boss_data, boss_num)
	max_boss_id := DB_get_maxBoss_id()
	rand.Seed(time.Now().Unix())
	for key, _ := range B_D {
		B_D[key].Boss_id = key + 1 + max_boss_id
		B_D[key].Hp = rand.Intn(1000) + rand.Intn(100000)*(rand.Intn(2)/2) + rand.Intn(1000000)*(rand.Intn(2)/2)
		B_D[key].Max_Hp = B_D[key].Hp
		B_D[key].Mola = B_D[key].Hp + rand.Intn(1000)
		B_D[key].Play_num = 0
	}
	DB_insert_Boss_data(B_D)
	c.Redirect(http.StatusMovedPermanently, "admpage")
}
func Boss_data_page(c *gin.Context) {
	B_D, _ := DB_get_Boss_Data()
	data := Get_cookie(c)
	data["Boss_Data"] = B_D
	c.HTML(http.StatusOK, "boss_page.html", data)
}
