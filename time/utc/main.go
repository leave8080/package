package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// 获取当前(当地)时间
	t := time.Now()
	// 获取0时区时间
	t = time.Now().UTC()
	fmt.Println(t)
	// 获取当前时间戳
	timestamp := t.Unix()
	fmt.Println(timestamp)
	// 获取时区信息
	name, offset := t.Zone()
	fmt.Println(name, offset)

	// 把时间戳转换为时间
	// 格式化时间
	sample := "2006-01-02 15:04:05"
	// 服务器时区
	fmt.Println("Current time 1 : ", time.Unix(timestamp+int64(offset), 0).Format(sample))
	// 零时区
	fmt.Println("Current time 2 : ", time.Unix(timestamp, 0).UTC().Format(sample))
	// 指定时间差
	fmt.Println("Current time 3 : ", time.Unix(timestamp-8*3600, 0).UTC().Format(sample))
	// 或
	fmt.Println("Current time 4 : ", time.Unix(timestamp, int64(8*time.Hour)).UTC().Format(sample))
	// 或
	fmt.Println("Current time 5 : ", time.Unix(timestamp, 0).UTC().Add(8*time.Hour).Format(sample))

	stamp, _ := time.ParseInLocation(sample, time.Unix(timestamp-8*3600, 0).UTC().Format(sample), time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	log.Println(stamp.Unix())
}
