package main

import "fmt"

func B(a, b int) int {
	return a + b
}
func A(f func(int, int) int) int {
	value := f(19, 19)
	return value
}
func main() {
	fmt.Println(A(B))
}
