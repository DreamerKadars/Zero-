package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Quick_Join_And_Hit_Html(c *gin.Context) {
	//拿到cookie的uid
	data := Get_cookie(c)
	uid, _ := strconv.Atoi(data["uid"].(string))
	U_D, err := DB_get_User_data(uid)
	if err != nil {
		fmt.Println("寻找User失败！！")
		fmt.Println(err.Error())
		return
	}
	//选择一个血量最少的boss
	B_D, _ := DB_Find_Min_Hp_Boss(1)
	if B_D == nil {
		fmt.Println("寻找Boss失败！！")
		return
	}

	Boss_id := B_D[0].Boss_id
	Mola := B_D[0].Mola
	Atk := U_D.Atk

	_, result := Hit_Boss(uid, Boss_id, Atk, Mola)

	data["result"] = result

	data["nextpage_flag"] = true
	data["nextpage"] = "index"

	c.HTML(http.StatusOK, "verify.html", data)
}

const Quick_Boss_num int = 5

func Quick_Join_And_Hit_Shell(c *gin.Context) {

	data := gin.H{}
	var result string
	//拿到post的uid和pwd，验证
	c.Request.ParseForm()
	var pwd = c.Request.Form["pwd"][0]
	uid_str := c.Request.Form["uid"][0]
	uid, _ := strconv.Atoi(uid_str)

	//将获取用户和，获取Boss信息并发起来
	var Err_chan chan error = make(chan error)
	var U_D User_data
	var B_D []Boss_data = make([]Boss_data, Quick_Boss_num)

	startTime := time.Now().UnixNano() //并发计时器
	go func() {
		//验证用户身份
		startTime := time.Now().UnixNano()
		var err error
		err = DB_found(uid, pwd)
		endTime := time.Now().UnixNano()
		seconds := float64((float64(endTime) - float64(startTime)) / 1e9)
		fmt.Println("验证用户信息用时：", seconds)
		Err_chan <- err
	}()
	go func(U_D *User_data) {
		//查用户相关
		startTime := time.Now().UnixNano()
		var err error
		*U_D, err = DB_get_User_data(uid)
		endTime := time.Now().UnixNano()
		seconds := float64((float64(endTime) - float64(startTime)) / 1e9)
		fmt.Println("寻找用户信息用时：", seconds)
		Err_chan <- err
	}(&U_D)

	go func(B_D *[]Boss_data) { //切片传的其实是个指针。。。，这个和数组区别挺大的
		//查boss相关
		startTime := time.Now().UnixNano()
		var err error
		//当前这个比较慢，居然是因为用了？，我表示非常不理解
		*B_D, err = DB_Find_Min_Hp_Boss(Quick_Boss_num)
		endTime := time.Now().UnixNano()
		seconds := float64((float64(endTime) - float64(startTime)) / 1e9)
		fmt.Println("寻找Boss信息用时：", seconds)
		Err_chan <- err
	}(&B_D)

	//从信道中取结果
	err1 := <-Err_chan
	err2 := <-Err_chan
	err3 := <-Err_chan
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("失败！！")
		return
	}
	close(Err_chan)

	//并发计时器结尾
	endTime := time.Now().UnixNano()
	seconds := float64((float64(endTime) - float64(startTime)) / 1e9)
	fmt.Println("将用户和找boss信息并发后时间为：", seconds)

	//随机选择一个Hit
	rand.Seed(time.Now().Unix())
	var boss_flag int = rand.Intn(Quick_Boss_num)

	_, result = Hit_Boss(uid, B_D[boss_flag].Boss_id, U_D.Atk, B_D[boss_flag].Mola)
	//等待结果

	data["result"] = result
	c.HTML(http.StatusOK, "Simple_Replay.html", data)
}
