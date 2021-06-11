package main

import "fmt"

func main() {

	ss := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(ss); i += 2 {
		fmt.Println(ss[i:Ac(i, len(ss))])
	}
}

func Ac(i, num int) int {
	if i+2 > num {
		return num
	}
	return i + 2
}
