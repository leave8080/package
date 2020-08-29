package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sqlx.DB
)

type MdMbooks struct {
	MemberId int    `db:"member_id"`
	Account  string `db:"account"`
	NickName string `db:"nickname"`
	Password string `db:"password"`
}

func initdb() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/mbook?charset=utf8mb4&parseTime=True"
	if DB, err = sqlx.Connect("mysql", dsn); err != nil {
		panic(err.Error())
	}

}

func main() {
	initdb()
	var mm []MdMbooks
	// select 获取多条记录， get 则是一条记录。
	if err := DB.Select(&mm, "select member_id,account,nickname,password from md_members "); err != nil {
		fmt.Println("select err:", err.Error())
		return
	}
	fmt.Println("mm", mm)
	for i, _ := range mm {
		fmt.Println(mm[i])
	}

	//
	sqlx.In("select * from md_members where member_id in (?)", "1 ")
}
