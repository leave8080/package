package main

import "fmt"

// 0 1 2 3 4 5 6 7
// 0 1 1 2 3 5 8
func fib1(num int) int {
	if num < 2 {
		return num
	}
	return fib1(num-2) + fib1(num-1)
}

func fib2(num int) int {
	if num < 2 {
		return num
	}
	first := 0
	second := 1
	for i := 0; i < num-1; i++ {
		sum := first + second
		first = second
		second = sum
	}
	return second
}
func main() {
	num := 60
	//fmt.Println(fib1(num))
	fmt.Println(fib2(num))
}
