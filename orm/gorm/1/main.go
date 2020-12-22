package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init1() (db *gorm.DB) {
	fmt.Println("mysql")
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/mbook?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		defer db.Close()
		logs.Error(err)
		return
	}

	return db
}

func main() {
	db = Init1()
	var count int

	if err := db.Table("md_category").Select("count(distinct(pid))").Count(&count).Error; err != nil {
		logs.Error(err)
		return
	}
	fmt.Println("num:", count)
	defer db.Close()
}
