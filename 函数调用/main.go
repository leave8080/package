package main

import "fmt"

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
func main() {
	fmt.Println(add(1, 2, 3))
	fmt.Println(add([]int{1, 2, 3}...))
}
