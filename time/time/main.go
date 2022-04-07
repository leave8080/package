package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime1 := time.Now()
	oldTime := currentTime1.AddDate(0, 0, -7)
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

	fmt.Println("#######")

	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(loc.String())

	data := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
	fmt.Println(data)

	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())

	fmt.Println(startTime)
	fmt.Println(startTime.Format("2006/01/02 15:04:05"))
	fmt.Println(endTime.Unix())

	fmt.Println(DayTime())

	fmt.Println(TimeStringToUnix("2022-03-23"))
}

func DayTime() int64 {
	currentTime := time.Now()

	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()-6, 0, 0, 0, 0, currentTime.Location())

	//fmt.Println(startTime)
	//
	//fmt.Println(startTime.Format("2006/01/02 15:04:05"))

	return startTime.Unix()
}

func TimeStringToUnix(timeSting string) int64 {

	loc, _ := time.LoadLocation("Asia/Shanghai")

	the_time, err := time.ParseInLocation("2006-01-02", timeSting, loc)

	if err == nil {

		return the_time.Unix() //1504082441

	}
	return 0

}
