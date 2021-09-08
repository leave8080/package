package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//t := "2019-10-10 10:10:10"
	//t1, _ := time.Parse("2006-01-02 15:04:05", t)
	//t2, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	//fmt.Println(t1)
	//fmt.Println(t2)
	//fmt.Println(t1.Equal(t2))
	//Asia/Tokyo

	var cstSh, _ = time.LoadLocation("America/Chicago") //上海
	fmt.Println("SH : ", time.Now().In(cstSh).Format("2006-01-02 15:04:05"))
	timeTemplate1 := "2006-01-02 15:04:05" //常规类型

	stamp, _ := time.ParseInLocation(timeTemplate1, time.Now().In(cstSh).Format("2006-01-02 15:04:05"), time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	log.Println(stamp.Unix())                                                                                       //输出：1546926630
	////时区转换
	//fmt.Println("***************")
	//t = "2021-01-11T23:46:05Z"
	//t1, _ = time.Parse("2006-01-02T15:04:05Z", t)
	//fmt.Println(t)
	//fmt.Println("SH : ", t1.In(cstSh).Format("2006-01-02 15:04:05"))

	sample := "2006-01-02 15:04:05"
	stamp1, _ := time.ParseInLocation(sample, time.Unix(time.Now().UTC().Unix()+(-18000), 0).UTC().Format(sample), time.Local)
	log.Println(stamp1.Unix())
}
