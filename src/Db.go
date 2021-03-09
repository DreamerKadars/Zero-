package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

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
	Uid     int `db:"uid"`
	Boss_id int `db:"Boss_id"`
	Hp      int `db:"Hp"`
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
		return
	}
	fmt.Println("连接数据库成功！")
	DB = database
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

//注册一个用户
func DB_register(uid int, pwd string) error {

	sql := "insert into User (uid,pwd) values (?,?)"
	r, err := DB.Exec(sql, uid, pwd)
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}
	fmt.Println("insert succ", id)
	fmt.Println(uid, " ", pwd)
	return nil
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
func DB_insert_User_data(U_D User_data) error {
	sql := "insert into User_data (uid,name,atk,mola,buff1,buff2,buff3) values (?,?,?,?,?,?,?)"
	r, err := DB.Exec(sql, U_D.Uid, U_D.Name, U_D.Atk, U_D.Mola, U_D.Buff1, U_D.Buff2, U_D.Buff3)
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,", err)
		return err
	}
	fmt.Println("insert succ", id)
	return nil
}

//查询展示Boos的信息
func DB_get_Boos_Data() ([]Boss_data, error) {
	var B_d []Boss_data
	sql := "select * from Boss_data "
	err := DB.Select(&B_d, sql)
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
