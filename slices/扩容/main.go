package main

import "fmt"

func main() {
	s := make([]int, 100, 100)
	fmt.Println(cap(s), len(s))
	s = append(s, 1)
	fmt.Println(cap(s))
}
