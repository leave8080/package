package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	Max := 6
	total := 10

	num, _ := strconv.ParseFloat(fmt.Sprintf("%f", float64(total)/float64(Max)), 64) // 保留2位小数
	page := math.Ceil(num)
	fmt.Println(page)
	for i := 0; i < int(page); i++ {
		fmt.Println(i*Max, Max)
	}
}
