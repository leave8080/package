package main

import (
	"fmt"
	"github.com/dariubs/percent"
	"strconv"
)

func main() {
	fmt.Println(percent.Percent(25, 200))

	num, _ := strconv.ParseFloat(fmt.Sprintf("%0.2f", float64(4)/float64(6)), 64) // 保留2位小数
	//page := math.Ceil(num)
	fmt.Println(num)

	//now := time.Now().Format("2006-01-02")

}
