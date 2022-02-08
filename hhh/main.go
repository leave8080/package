package main

import (
	"log"
	"time"
)

func main() {
	sample := "2006-01-02 15:04:05"
	stamp, _ := time.ParseInLocation(sample, time.Unix(time.Now().UTC().Unix()+int64(32400), 0).UTC().Format(sample), time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	log.Println(stamp.Unix())

	log.Println(int32(time.Now().Unix() + int64(32400)))

}
