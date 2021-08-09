package main

import (
	"fmt"
	"gofer/pkg/log"
	"strconv"
	"time"
)

func main() {
	s := "大"
	fmt.Println(len(s))

	lon := strconv.FormatFloat(23.125, 'W', 6, 64)
	lat := strconv.FormatFloat(113.361, 'f', 6, 64)

	fmt.Println(lon, lat)

	//ss := "12-28"

	//如果今天的月>2
	nTime := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	log.Info(nTime)
	log.Info(time.Now())
}
