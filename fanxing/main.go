package main

import "fmt"

func printSlice[T any](s []T) {
	for _, V := range s {
		fmt.Printf("%v \n", V)
	}
}

func main() {
	printSlice[int]([]int{66, 77, 88, 99, 100})
	printSlice[string]([]string{"zhangsan", "lisi", "wangwu", "zhaosi"})
	printSlice([]int{66, 77, 88, 99, 100})
}
