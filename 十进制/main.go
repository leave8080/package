package main

import (
	"fmt"
	"strconv"
)

// 将十进制数字转化为二进制字符串
func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}

func main() {
	fmt.Println(
		convertToBin(2),
		convertToBin(19),
		convertToBin(15),
		convertToBin(0),
	)
}
