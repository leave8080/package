package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type ExampleConf struct {
	ID       int64  `orm:"column(id);pk;auto" json:"id"`       //主键
	Driver   string `orm:"column(driver)" json:"driver"`       //采集服务名称
	Model    string `orm:"column(model);null" json:"model"`    //保留字段
	ExScope  string `orm:"column(ex_scope)" json:"exScope"`    //范围
	Name     string `orm:"column(name)" json:"name"`           //字段中文名称
	Key      string `orm:"column(key)" json:"key"`             //字段英文名称
	Value    string `orm:"column(value)" json:"value"`         //字段默认值
	Ed       string `orm:"column(ed)" json:"ed"`               //描述
	Flag     string `orm:"column(flag)" json:"flag"`           //同步备注
	TextType string `orm:"column(text_type)" json:"textType"`  //字段类型
	TextData string `orm:"column(text_data)" json:"textData"`  //字段类型默认值
	Data1    string `orm:"column(re_data1);null" json:"data1"` //保留字段1
	Data2    string `orm:"column(re_data2);null" json:"data2"` //保留字段2
}

func main() {
	db, err := sql.Open("sqlite3", "./data12.db")
	if err != nil {
		log.Println("open sql err", err)
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("update upoint set gn=?||pn")

	res, err := stmt.Exec(`W3.DATAPOINT.`)

	affect, err := res.RowsAffected()

	fmt.Println(affect)

}
