package main

import "fmt"

func main() {
	s := 123456789
	fmt.Println(NewNum(s, 3))
}

func NewNum(s int, n int) int {

	num := s / (1000)
	return num
}
