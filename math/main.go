package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	pagesiez := 20000
	//num := 21 / pagesiez
	num1, _ := strconv.ParseFloat(fmt.Sprintf("%f", float64(9)/float64(pagesiez)), 64) // 保留2位小数
	fmt.Println(num1)
	//fmt.Printf("[3.14]向上取整为:[%.f]\n", math.Ceil(3.14))
	page := math.Ceil(num1)
	//math.Ceil()
	//fmt.Println(math.Ceil(float64(2.1)))
	//fmt.Println(page)
	for i := 0; i < int(page); i++ {
		fmt.Println(i*pagesiez, pagesiez)
	}

}
