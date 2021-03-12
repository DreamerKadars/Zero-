package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Get_Boss_data_str(B_D Boss_data) []string {
	var Boss_data_str []string = make([]string, 5)
	Boss_data_str[0] = "Boss id    :" + strconv.Itoa(B_D.Boss_id)
	Boss_data_str[1] = "剩余HP     :" + strconv.Itoa(B_D.Hp)
	Boss_data_str[2] = "最大HP     :" + strconv.Itoa(B_D.Max_Hp)
	Boss_data_str[3] = "竞争玩家数量:" + strconv.Itoa(B_D.Play_num)
	Boss_data_str[4] = "赏金数额    :" + strconv.Itoa(B_D.Mola)
	return Boss_data_str
}
func Creat_Boss_rand(num int) []Boss_data {
	max_boss_id := DB_get_maxBoss_id()
	var B_D []Boss_data = make([]Boss_data, num)
	rand.Seed(time.Now().Unix())
	for key, _ := range B_D {
		B_D[key].Boss_id = key + 1 + max_boss_id
		B_D[key].Hp = rand.Intn(1000) + rand.Intn(100000)*(rand.Intn(2)/2) + rand.Intn(1000000)*(rand.Intn(2)/2)
		B_D[key].Max_Hp = B_D[key].Hp
		B_D[key].Mola = B_D[key].Hp + rand.Intn(1000)
		B_D[key].Play_num = 0
	}
	return B_D
}
func Boss_data_add(c *gin.Context) {
	c.Request.ParseForm()
	boss_num, _ := strconv.Atoi(c.Request.Form["boss_num"][0])

	B_D := Creat_Boss_rand(boss_num)
	DB_insert_Boss_data(B_D)

	c.Redirect(http.StatusMovedPermanently, "admpage")
}
func Boss_data_page(c *gin.Context) {
	B_D, _ := DB_get_Boss_Data_Live()
	data := Get_cookie(c)
	data["Boss_Data"] = B_D
	c.HTML(http.StatusOK, "boss_page.html", data)
}
