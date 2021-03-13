package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

const DB_extra_num int = 4

var DB_extra []*sqlx.DB = make([]*sqlx.DB, DB_extra_num)

type User struct {
	Uid int    `db:"uid"`
	Pwd string `db:"pwd"`
}
type User_data struct {
	Uid   int    `db:"uid"`
	Name  string `db:"name"`
	Atk   int    `db:"atk"`
	Mola  int    `db:"mola"`
	Buff1 int    `db:"buff1"`
	Buff2 int    `db:"buff2"`
	Buff3 int    `db:"buff3"`
}
type Boss_data struct {
	Boss_id  int `db:"Boss_id"`
	Hp       int `db:"Hp"`
	Max_Hp   int `db:"Max_Hp"`
	Play_num int `db:"play_num"`
	Mola     int `db:"mola"`
}
type User_history struct {
	Uid     int  `db:"uid"`
	Boss_id int  `db:"Boss_id"`
	Hp      int  `db:"Hp"`
	IsKill  bool `db:"IsKill"`
}
type Now_Battle struct {
	Boss_id int `db:"Boss_id"`
	Uid     int `db:"uid"`
}
type Adm struct {
	Uid int `db:"uid"`
}

//初始化连接数据库
func init() {
	fmt.Print("开始连接数据库。。。")
	database, _ := sqlx.Open("mysql", "root:ma794866734@tcp(81.70.159.251:3306)/Zero")
	err := database.Ping()
	if err != nil {
		fmt.Println("连接数据库失败！")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("连接数据库成功！")
	fmt.Println("建立信道，接受用户请求")
	go DB_deal_hit()

	DB = database

	for key := range DB_extra {
		DB_extra[key], _ = sqlx.Open("mysql", "root:ma794866734@tcp(81.70.159.251:3306)/Zero")
	}
}

//测试连接
func DB_connect() error {

	database, _ := sqlx.Open("mysql", "root:ma794866734@tcp(81.70.159.251:3306)/Zero")
	err := database.Ping()
	if err != nil {
		fmt.Println("连接数据库失败！")
		return errors.New("连接失败")
	}
	fmt.Println("连接数据库成功！")
	DB = database
	return nil
}

//注册一个用户,指定数据库指针
func DB_register_point(uid int, pwd string, DB_p *sqlx.DB) error {

	sql := "insert into User (uid,pwd) values (?,?)"
	_, err := DB_p.Exec(sql, uid, pwd)
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}
	//fmt.Println("成功插入用户：", uid)
	return nil
}

//注册很多个用户
func DB_add_user(user_num int) int {
	fmt.Println("开始插入用户。。。。")
	startTime := time.Now().UnixNano()
	var affect int = 0
	var max_uid []sql.NullInt32
	sql := "select max(uid) as uid from User"
	err := DB.Select(&max_uid, sql)

	if err != nil {
		fmt.Println(err.Error())
		return affect
	}

	var now_id int
	if max_uid[0].Valid {
		now_id = int(max_uid[0].Int32)
		now_id = now_id + 1
	} else {
		now_id = 1
	}
	var Last chan error = make(chan error)
	var con_num int = 0
	var dis_con_num = 0
	var Binfa_num int = 100

	var con_num_chan chan bool = make(chan bool)     //处理连接信号信道
	var dis_con_num_chan chan bool = make(chan bool) //处理断开连接信号信道

	go func() { //处理接入请求
		for range con_num_chan {
			if con_num-dis_con_num > Binfa_num {
				time.Sleep(time.Duration(1) * time.Second)
				fmt.Println("并发限制，等待！", con_num-dis_con_num)
			}
			fmt.Println("接入连接")
			con_num++
		}
	}()
	go func() { //处理断开请求
		for range dis_con_num_chan {
			dis_con_num++
			fmt.Println("断开连接")
		}
	}()
	var E []error = make([]error, 0)      //汇总所有连接的结束状态
	var ans_num chan int = make(chan int) //汇总最后的数量
	go func() {                           //处理结束请求
		for i := now_id; i < now_id+user_num; i++ {
			err = <-Last
			if err == nil {
				affect++
			} else {
				E = append(E, err)
			}
		}
		close(con_num_chan) //全部接受完毕，关闭信道
		for _, value := range E {
			println(value.Error())
		}
		endTime := time.Now().UnixNano()
		seconds := float64((float64(endTime) - float64(startTime)) / 1e9)
		fmt.Println("总共完成：", affect)
		fmt.Println("总共用时：", seconds, "s")
		fmt.Println("平均用时：", seconds/float64(affect), "s")
		ans_num <- affect
	}()

	for i := now_id; i < now_id+user_num; i++ {
		go func(uid int) {
			var try_time int = 0
			const max_try_time = 10 //最大尝试次数
			con_num_chan <- true
			for {
				E := DB_register_point(uid, strconv.Itoa(uid), DB_extra[uid%DB_extra_num])
				if E == nil {
					break
				} else {
					fmt.Println(E.Error())
				}
				fmt.Println(E.Error())
				time.Sleep(time.Duration(1) * time.Second)
				try_time++
				fmt.Println("失败，继续尝试！")
				if try_time > max_try_time {
					fmt.Println("彻底失败，放弃尝试！")
					return
				}
			}
			var U_D User_data = User_data{uid, "normal", 35, 0, 0, 0, 0}
			for {
				E := DB_insert_User_data_point(U_D, DB_extra[uid%DB_extra_num])
				if E == nil {
					break
				} else {
					fmt.Println(E.Error())
				}
				time.Sleep(time.Duration(1) * time.Second)
				try_time++
				fmt.Println("失败，继续尝试！")
				if try_time > max_try_time {
					fmt.Println("彻底失败，放弃尝试！")
					return
				}
			}
			Last <- err
			dis_con_num_chan <- true
		}(i)
	}

	return <-ans_num
}

//一个用户是否存在
func DB_found(uid int, pwd string) error {
	var usr []User
	sql := "select uid, pwd from User where uid=? && pwd =? "
	err := DB.Select(&usr, sql, uid, pwd)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return err
	}
	if len(usr) == 0 {
		return errors.New("用户名或者密码错误！")
	}
	return nil
}

//向指定目录中输出num个uid,pwd
func DB_Get_User_file(num int, filename string) error {
	var usr []User
	sql := "select * from User limit ?"
	err := DB.Select(&usr, sql, num)

	if err != nil {
		fmt.Println("exec failed, ", err)
		return err
	}
	f, err1 := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	if err1 != nil {
		fmt.Println("exec failed, ", err1)
		return err1
	}
	for key := range usr {
		f.WriteString("uid=" + strconv.Itoa(usr[key].Uid) + "&pwd=" + usr[key].Pwd + "\n")
	}

	f.Sync()
	fmt.Println("写完了")
	return nil
}

//检查一个用户是否是adm
func DB_is_adm(uid int) error {
	var adm []Adm
	sql := "select uid from Adm where uid=?"
	err := DB.Select(&adm, sql, uid)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return err
	}
	if len(adm) == 0 {
		return errors.New("不是管理员！")
	}
	return nil
}

var DB_name = [6]string{"User", "User_data", "Boss_data", "User_history", "Now_Battle", "Adm"}

// <p>用户登陆表User: </p>
//     <p>用户信息表User_data: </p>
//     <p>Boss信息表Boss_data: </p>
//     <p>用户历史User_history: </p>
//     <p>对战信息Now_Battle: </p>
//     <p>管理员账号Adm: </p>
//查找指定表的数量
func DB_num(table_name string) int {
	sql := "select count(*) from " + table_name
	var num []int
	DB.Select(&num, sql)
	return num[0]
}

func DB_get_User_data(uid int) (User_data, error) {
	var U_d []User_data
	sql := "select * from User_data where uid=?"
	err := DB.Select(&U_d, sql, uid)
	if err != nil {
		fmt.Println("exec failed, ", err)
		var err_data User_data
		return err_data, err
	}
	if U_d == nil {
		fmt.Println("exec failed, 没有找到用户信息")
		var err_data User_data
		return err_data, errors.New("没有找到该用户信息")
	}
	return U_d[0], nil
}

//为用户注册信息
func DB_insert_User_data_point(U_D User_data, DB_p *sqlx.DB) error {
	sql := "insert into User_data (uid,name,atk,mola,buff1,buff2,buff3) values (?,?,?,?,?,?,?)"
	_, err := DB_p.Exec(sql, U_D.Uid, U_D.Name, U_D.Atk, U_D.Mola, U_D.Buff1, U_D.Buff2, U_D.Buff3)
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}
	//fmt.Println("成功插入用户数据！")
	return nil
}

//查询展示Boos的信息
func DB_get_Boss_Data(limit string) ([]Boss_data, error) {
	var B_d []Boss_data
	sql := "select * from Boss_data " + limit
	err := DB.Select(&B_d, sql)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return B_d, err
	}
	return B_d, nil
}

//根据生存分类
func DB_get_Boss_Data_Live() ([]Boss_data, error) {
	return DB_get_Boss_Data("where Hp != 0 ")
}
func DB_get_Boss_Data_Die() ([]Boss_data, error) {
	return DB_get_Boss_Data("where Hp = 0 ")
}

//得到存活的Boss num
func DB_get_Live_Boss_num() int {
	sql_str := "select count(*) from Boss_data where Hp>0"
	var num []sql.NullInt32
	DB.Select(&num, sql_str)
	if num == nil {
		return 10000000
	}
	if num[0].Valid {
		return int(num[0].Int32)
	} else {
		return 0
	}
}

//得到一个boss信息
func DB_get_Boss_Data_one(boss_id int) ([]Boss_data, error) {
	var B_d []Boss_data
	sql := "select * from Boss_data where boss_id = ?"
	err := DB.Select(&B_d, sql, boss_id)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return B_d, err
	}
	return B_d, nil
}

//插入Boss信息
//为Boss注册信息
func DB_insert_Boss_data(B_d []Boss_data) error {
	if B_d == nil {
		return errors.New("要插入的内容为空！")
	}
	sql := "insert into Boss_data (Boss_id,Hp,Max_HP,play_num,mola) values (?,?,?,?,?)"
	for _, Boss := range B_d {
		_, err := DB.Exec(sql, Boss.Boss_id, Boss.Hp, Boss.Max_Hp, Boss.Play_num, Boss.Mola)
		if err != nil {
			fmt.Println("exec failed,", err)
			return err
		}
	}
	return nil
}

//求Boss_id的最大值
func DB_get_maxBoss_id() int {
	sql_str := "select max(Boss_id) from Boss_data"
	var num []sql.NullInt32
	DB.Select(&num, sql_str)
	if num[0].Valid {
		return int(num[0].Int32)
	} else {
		return 0
	}
}

//获得用户当前正在进行的Battle
func DB_get_Battle(uid int) []Now_Battle {
	var now_battle []Now_Battle
	sql_str := "select * from Now_Battle where uid = ?"
	DB.Select(&now_battle, sql_str, uid)
	return now_battle
}

//获得用户过去的历史--详细
func DB_get_History_all(uid int) []User_history {
	var user_history []User_history
	sql_str := "select * from User_history where uid = ?"
	DB.Select(&user_history, sql_str, uid)
	return user_history
}

//获得用户过去的历史--粗略
func DB_get_History_group(uid int) []User_history {
	var user_history []User_history
	sql_str := "select Boss_id,sum(Hp) as Hp,sum(IsKill) as IsKill from User_history where uid = ? group by Boss_id;"
	DB.Select(&user_history, sql_str, uid)
	return user_history
}

//用户参与战斗
func DB_join_battle(uid, Boss_id int) error {
	a, _ := DB_get_Boss_Data_one(Boss_id)
	if a == nil {
		return errors.New("此id不存在！！！")
	}
	if a[0].Hp == 0 {
		return errors.New("此Boss已经被击败！！！")
	}

	sql := "insert into Now_Battle (Boss_id,uid) values (?,?)"
	r, err := DB.Exec(sql, Boss_id, uid)
	if err != nil {

		if err.Error()[12] == 'D' && err.Error()[13] == 'u' {
			return errors.New("你已经参与了此场战斗！")
		}
		fmt.Println("exec failed,", err)
		return err
	}
	//add——playnum
	sql = "update Boss_data set play_num=play_num+1 where Boss_id=?"
	_, err = DB.Exec(sql, Boss_id)
	if err != nil {

		fmt.Println(err.Error())
	}
	_, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}
	return nil
}
func DB_exit_battle(uid, Boss_id int) error {
	a, _ := DB_get_Boss_Data_one(Boss_id)
	if a == nil {
		return errors.New("此id不存在！！！")
	}
	sql := "delete from Now_Battle where uid=? and Boss_id=?"
	_, err := DB.Exec(sql, uid, Boss_id)
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}
	sql = "update Boss_data set play_num=play_num-1 where Boss_id=?"
	_, err = DB.Exec(sql, Boss_id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

//与该用户打同一boss的对手
func DB_Compete(uid int, boss_id int) []User_data {
	var user_data []User_data
	sql_str := "select Now_battle.uid,name from Now_battle inner join User_data on Now_battle.uid=User_data.uid where User_data.uid != ? and boss_id=?"
	DB.Select(&user_data, sql_str, uid, boss_id)
	return user_data
}

var M map[int]sync.Mutex = make(map[int]sync.Mutex)

//打击信道结构
type Hit struct {
	uid     int
	boss_id int
	atk     int
	Re_chan chan int
}

var Hit_ch chan Hit = make(chan Hit)

//处理打击
func DB_deal_hit() {
	for v := range Hit_ch {
		err := DB_Hit_Boss(v) //通过信号的方式传递信息，自动阻塞
		if err != nil {
			//暂时屏蔽，因为无效错误不好观察
			//fmt.Println(err.Error())
		}
	}
}

//进行打击
func DB_Hit_Boss(h Hit) error { //对于同一个boss，要求只有一个用户能进入请求
	//查这个boss还有多少hp
	sql_str := "select Hp from Boss_data where Boss_id = ?"
	var num []sql.NullInt32
	DB.Select(&num, sql_str, h.boss_id) //拿boss数据
	if num == nil || !num[0].Valid {
		h.Re_chan <- -1
		close(h.Re_chan)
		return errors.New("没有这个Boss")
	}
	Hp := num[0].Int32
	if Hp == 0 {
		h.Re_chan <- -1
		close(h.Re_chan)
		return errors.New("此Boss已经挂了")
	}
	if Hp-int32(h.atk) > 0 {
		Hp = Hp - int32(h.atk)
	} else {
		Hp = 0
	}
	sql_str = "update Boss_data set HP=? where Boss_id=?"
	_, err := DB.Exec(sql_str, Hp, h.boss_id) //更新boss状态
	if Hp == 0 {
		sql_str := "select mola from Boss_data where Boss_id = ?" //看它值多少
		var num []sql.NullInt32
		err = DB.Select(&num, sql_str, h.boss_id)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		Mola_get := num[0].Int32
		sql_str = "update User_data set mola=mola+? where uid=?" //加钱
		_, err := DB.Exec(sql_str, Mola_get, h.uid)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if err != nil {
		h.Re_chan <- -1
		close(h.Re_chan)
		return errors.New("数据库错误")
	}
	sql_str = "insert into User_history (uid,Boss_id,IsKill,Hp) values (?,?,?,?)" //用户记录
	var IsKill int
	if Hp == 0 {
		IsKill = 1
	} else {
		IsKill = 0
	}
	_, err = DB.Exec(sql_str, h.uid, h.boss_id, IsKill, h.atk) //更新用户历史
	if err != nil {
		h.Re_chan <- -1
		close(h.Re_chan)
		return errors.New("数据库错误")
	}
	if Hp == 0 {
		h.Re_chan <- 0 //击杀
	} else {
		h.Re_chan <- int(Hp) //剩余血量
	}

	close(h.Re_chan)
	return nil

	//够不够？
	//改一下，是否改成功？
}

//找到血量最少且存活的那些Boss
func DB_Find_Min_Hp_Boss(Boss_num int) ([]Boss_data, error) {
	var B_d []Boss_data
	sql := "select * from Boss_data where Hp>0 order by Hp limit ?"
	err := DB.Select(&B_d, sql, Boss_num)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return B_d, err
	}
	return B_d, nil
}

type Adm_data struct {
	Table_Name string  `db:"表名"`
	Log_num    int     `db:"记录数"`
	Data_M     float32 `db:"数据容量(MB)"`
	Index_M    float32 `db:"索引容量(MB)"`
}

//拿到数据库管理信息
func Get_Adm_Data() ([]Adm_data, error) {
	var A_d []Adm_data
	sql := `select
	table_name as '表名',
	table_rows as '记录数',
	truncate(data_length/1024/1024, 2) as '数据容量(MB)',
	truncate(index_length/1024/1024, 2) as '索引容量(MB)'
	from information_schema.tables
	where table_schema ='Zero'
	order by data_length desc, index_length desc;`

	err := DB.Select(&A_d, sql)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return A_d, err
	}
	return A_d, nil
}
