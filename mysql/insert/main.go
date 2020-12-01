package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

var Db *sqlx.DB

//与数据库建立连接
func init() {
	database, err := sqlx.Open("mysql", "root:123@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
}

func main() { //事务操作
	conn, err := Db.Begin()
	if err != nil {
		return
	}
	//插入数据
	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err) //回滚　　　　 conn.Rollback()
		return
	}

	fmt.Println("insert succ:", id) //提交事务
	conn.Commit()
}
