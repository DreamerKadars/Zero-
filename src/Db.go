package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

type User struct {
	Uid int    `db:"uid"`
	Pwd string `db:"pwd"`
}

func init() {
	database, err := sqlx.Open("mysql", "root:ma794866734@tcp(81.70.159.251:3306)/Zero")
	if err != nil {
		fmt.Println("mysql数据库连接失败！", err)
		os.Exit(0)
		return
	}
	fmt.Println(database)
	fmt.Println("mysql数据库连接成功！")

	DB = database
}
func DB_register(uid int, pwd string) {
	sql := "insert into User (uid,pwd) values (?,?)"
	r, err := DB.Exec(sql, uid, pwd)
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,", err)
		return
	}
	fmt.Println("insert succ", id)
	fmt.Println(uid, " ", pwd)

}

// func register(uid int, pwd string) {
// 	sql := "insert into User (uid,pwd) values (?,?)"
// 	r, err := db.Exec(sql, uid, pwd)
// 	if err != nil {
// 		fmt.Println("exec failed,", err)
// 		return
// 	}
// 	id, err := r.LastInsertId()
// 	if err != nil {
// 		fmt.Println("exec failed,", err)
// 		return
// 	}
// 	fmt.Println("insert succ", id)
// 	fmt.Println(uid, " ", pwd)
// }
// func found(uid int, pwd string) bool {
// 	var usr []User
// 	sql := "select uid, pwd from User where uid=? && pwd =? "
// 	err := db.Select(&usr, sql, uid, pwd)
// 	if err != nil {
// 		fmt.Println("exec failed, ", err)
// 		return false
// 	}

// 	if len(usr) == 1 {
// 		fmt.Println("select succ:", usr)
// 		return true
// 	}
// 	return false
// }
