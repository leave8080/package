package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	currentTime := time.Now()
	m, _ := time.ParseDuration("-1h")

	result := currentTime.Add(m)

	fmt.Println(result.Format("2006-01-02 15:00:00"))

	fmt.Println(result.Format("2006/01/02 15:04:05"))
	fmt.Println(time.Now().Hour())
	var s string

	for i := 0; i < time.Now().Hour(); i++ {
		s += fmt.Sprintf(" SELECT %s HOUR UNION ALL ", strconv.Itoa(i))
	}
	fmt.Println(s)
}
