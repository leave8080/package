package main

import (
	"fmt"
	"time"
)

func main() {
	//1618761600
	//1618675200
	//stime := int64(1617984000) //4-10
	stime := int64(1618761600) //4-19
	Remind(stime, 2)
	NextRemind(stime, 2)
	localRemind(1618675200, 1618851600, 2)
}

func localRemind(stime, Now int64, rate int64) {
	t := (Now - stime) / (60 * 60 * 24)
	t = t / rate
	fmt.Println(stime + t*60*60*24*rate)

}

func Remind(stime int64, rate int64) {
	t := (time.Now().Unix() - stime)
	count := t / (60 * 60 * 24)
	count /= rate
	localtime := stime + count*60*60*24*rate
	fmt.Println(t)
	fmt.Println(count)
	fmt.Println(localtime)
}
func NextRemind(stime int64, rate int64) {
	t := (time.Now().Unix() - stime)
	count := t / (60 * 60 * 24)
	count /= rate
	nexttime := stime + count*60*60*24*rate + rate*60*60*24
	fmt.Println(nexttime)
}
func GetTimeArr(start, end int64) int64 {
	//timeLayout  := "2006/01/02"
	//loc, _ := time.LoadLocation("Local")
	// 转成时间戳
	startUnix := int64(1618761600)
	endUnix := end
	startTime := startUnix
	endTime := endUnix
	// 求相差天数
	date := (endTime - startTime) / 86400
	return date
}
