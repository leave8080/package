package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Unix() - 600)
	fmt.Println(string(byte(52)))
	fmt.Println(string(byte(19)))
	var i = make([]int, 0)
	fmt.Println(i)
}
