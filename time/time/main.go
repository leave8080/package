package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -7)
	fmt.Println(oldTime)

	fmt.Println(time.Now().Day())

	nowTime := time.Now()
	getTime := nowTime.AddDate(0, 0, -1)                //年，月，日   获取一天前的时间
	resTime := getTime.Format("2006-01-02 15:04:05+08") //获取的时间的格式
	fmt.Println(resTime)

	getTime = nowTime.AddDate(0, -1, 0)             //年，月，日   获取一个月前的时间
	resTime = getTime.Format("2006-01-02 15:04:05") //获取的时间的格式
	fmt.Println(resTime)

	getTime = nowTime.AddDate(-2, 0, 0)  //年，月，日   获取两年前的时间
	resTime = getTime.Format("20060102") //获取的时间的格式
	fmt.Println(resTime)

	fmt.Println(time.Now().UnixNano() / 1000 / 1000)

	fmt.Println(time.Now().Unix() + 60*60*24*30*2)

	tims := "2016-01-02 " + "03:04" + ":00"
	fmt.Println(tims)
	timeTemplate1 := "2006-01-02 15:04:05"
	stamp, _ := time.ParseInLocation(timeTemplate1, tims, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	fmt.Println(stamp.Unix())

}
