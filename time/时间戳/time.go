package main

import (
	"fmt"
	"time"
)

const formart = "2006-01-02 15:04:05"

func main() {

	ss := time.Now().Format("2006-01-02") + " 15:00:00"
	fmt.Println(ss)
	stamp, _ := time.ParseInLocation(formart, ss, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	fmt.Println("waeaea", stamp.Unix())                       //输出：1546926630
	test()
	test2()
}

func test() {
	now := time.Now()
	timestamp := now.Unix() - int64(now.Second()) - int64((60 * now.Minute()))
	fmt.Println(timestamp, time.Unix(timestamp, 0), now.Unix())
}

func test2() {
	fmt.Println(time.Now().Format("2006-01-02"))
}
