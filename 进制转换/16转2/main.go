package main

import (
	"fmt"
)

func main() {
	data := 10
	for i := 0; i < 16; i++ {

		fmt.Println(data & 1)
		data >>= 1
	}
	fmt.Println("@@@@@@@@@@")
	j := 0
	for num := 0xf; num&(num-1) != 0; num = num & (num - 1) {
		j++
		fmt.Println(j)
	}
	fmt.Println("dasda")
	fmt.Println(j)
	fmt.Println(0xf - 1)

	num := 0xf
	i := 0
	for num != 0 {
		num = num & (num - 1)
		i++
	}
	fmt.Println(i)
	fmt.Println("###")

}
