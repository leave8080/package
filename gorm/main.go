package main

import (
	//导入MYSQL数据库驱动，这里使用的是GORM库封装的MYSQL驱动导入包，实际上大家看源码就知道，这里等价于导入github.com/go-sql-driver/mysql
	//这里导入包使用了 _ 前缀代表仅仅是导入这个包，但是我们在代码里面不会直接使用。
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
//在这里User类型可以代表mysql users表
type User struct {
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	//这里golang定义的Username变量和MYSQL表字段username一样，他们的名字可以不一样。
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	CreateTime int64 `gorm:"column:createtime"`
}

//设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}

func main() {
	//配置MySQL连接参数
	username := "root"         //账号
	password := "mxcloud@2020" //密码
	host := "106.14.218.147"   //数据库地址，可以是Ip或者域名
	port := 3308               //数据库端口
	Dbname := "test"           //数据库名

	//通过前面的数据库参数，拼接MYSQL DSN， 其实就是数据库连接串（数据源名称）
	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	//类似{username}使用花括号包着的名字都是需要替换的参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	//连接MYSQL
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	//定义一个用户，并初始化数据
	u := User{
		Username:   "tizi365",
		Password:   "123456",
		CreateTime: time.Now().Unix(),
	}
	db.AutoMigrate(&User{})
	//插入一条用户数据
	//下面代码会自动生成SQL语句：INSERT INTO `users` (`username`,`password`,`createtime`) VALUES ('tizi365','123456','1540824823')
	if err := db.Create(u).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

	//查询并返回第一条数据
	//定义需要保存数据的struct变量
	u = User{}
	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'tizi365') LIMIT 1
	isNotFound := db.Where("username = ?", "tizi365").First(&u).RecordNotFound()
	if isNotFound {
		fmt.Println("找不到记录")
		return
	}
	//打印查询到的数据
	fmt.Println(u.Username, u.Password)

	//更新
	//自动生成Sql: UPDATE `users` SET `password` = '654321'  WHERE (username = 'tizi365')
	//db.Model(User{}).Where("username = ?", "tizi365").Update("password", "654321")

	//删除
	//自动生成Sql： DELETE FROM `users`  WHERE (username = 'tizi365')
	//db.Where("username = ?", "tizi365").Delete(User{})

	// 根据主键查询第一条记录
	db.First(&u)
	fmt.Println(u.Username, u.Password)
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	// 随机获取一条记录
	db.Take(&u)
	fmt.Println(u.Username, u.Password)
	//// SELECT * FROM users LIMIT 1;

	// 根据主键查询最后一条记录
	db.Last(&u)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Println(u.Username, u.Password)

	// 查询所有的记录
	db.Find(&u)
	//// SELECT * FROM users;

	fmt.Println(u.Username, u.Password)

	// 查询指定的某条记录(仅当主键为整型时可用)
	db.First(&u, 10)
	//// SELECT * FROM users WHERE id = 10;
	fmt.Println(u.Username, u.Password)
}
