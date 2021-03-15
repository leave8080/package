package main

import (
	"fmt"
	"strconv"
	"time"
)

func f2i(f float64) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", f))
	return i
}
func main() {
	fmt.Println(f2i(100.0))
	//tt := time.Now()
	//fmt.Println(tt.Unix() - int64(60*tt.Second()) - int64(60*tt.Minute()) - int64(24*tt.Hour()))
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	timeNumber := t.Unix()
	timeNumber += 6 * 60 * 60
	fmt.Println("timeNumber:", timeNumber)
	fmt.Println(fmt.Sprintf("%d分%d秒", 391/60, 391%60))
}
