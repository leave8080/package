package main

import (
	"fmt"
	"time"
)

func main() {
	tt := 491
	fmt.Println(tt / 60)
	fmt.Println(tt % 60)

	now := time.Now()
	endTime10 := now.Unix() - int64(now.Second()) - int64(60*now.Minute())
	beginTime10 := endTime10 - 3600
	endTime, beginTime := endTime10*1000, beginTime10*1000
	fmt.Println(endTime, beginTime)
}
