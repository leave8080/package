package main

import (
	"fmt"
	"time"
)

func main() {
	timeZone := 28800
	sample := "2006-01-02 15:04:05"
	stamp, _ := time.ParseInLocation(sample, time.Unix(time.Now().UTC().Unix()+int64(timeZone), 0).UTC().Format(sample), time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	fmt.Println(stamp)
	fmt.Println(stamp.Day())
	fmt.Println(time.Unix(stamp.Unix(), 0).Format("2006-01-02"))
}

//2021-11-15 19:35:17 +0800 CST
