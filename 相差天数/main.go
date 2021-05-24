package main

import (
	"fmt"
	"time"
)

func GetTimeArr(start, end string) int64 {
	timeLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	// 转成时间戳
	startUnix, _ := time.ParseInLocation(timeLayout, start, loc)
	endUnix, _ := time.ParseInLocation(timeLayout, end, loc)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	// 求相差天数
	date := (endTime - startTime) / 86400
	return date
}

func main() {
	fmt.Println(GetTimeArr("2021-03-15", "2021-03-12"))
	fmt.Println(time.Now().AddDate(0, 0, -3).Format("2006-01-02"))
}
