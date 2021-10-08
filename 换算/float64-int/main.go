package main

import (
	"fmt"
	"strconv"
	"time"
)

func f2i(f int) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%02d", f))
	return i
}
func main() {
	fmt.Println(f2i(int(100.0)))
	//tt := time.Now()
	//fmt.Println(tt.Unix() - int64(60*tt.Second()) - int64(60*tt.Minute()) - int64(24*tt.Hour()))
	timeStr := time.Now().Format("2006-01-02")
	fmt.Println("timeStr:", timeStr)
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	timeNumber := t.Unix()
	timeNumber += 6 * 60 * 60
	fmt.Println("timeNumber:", timeNumber)
	fmt.Println(fmt.Sprintf("%d分%d秒", 391/60, 391%60))

	////
	catBathroomTimeStart := int(325)
	hour := f2i(catBathroomTimeStart) / 60
	second := f2i(catBathroomTimeStart) % 60
	fmt.Println(hour, second)

	fmt.Println(fmt.Sprintf("%02d", 5))

	ss := fmt.Sprintf("%0.1f", float32(17)/60)
	fs, _ := strconv.ParseFloat(ss, 64)
	fmt.Println(fs)
}
