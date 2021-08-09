package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Unix(time.Now().Unix(), 0))

	ret := []int{1, 2}
	fmt.Println(ret[0:0])
	ret = ret[0:0]
	fmt.Println(ret)

	ss := 1627520961
	timeStr := time.Unix(int64(ss), 0).Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	stime := t.Unix()
	fmt.Println(stime)

	sss := time.Now().Unix()
	r := time.Unix(sss, 0).Format("2006-01-02")
	fmt.Println(r)
}
