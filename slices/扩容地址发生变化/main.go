package main

import (
	"fmt"
)

func newslice(slice []int, num int) []int {
	return append(slice, num)
}
func main() {
	//var slice []int
	slice := make([]int, 0, 0)
	slice = append(slice, 1, 2, 3)
	fmt.Println(len(slice), cap(slice))
	nslice := newslice(slice, 4)
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(len(nslice), cap(nslice))
	fmt.Println(&nslice[0], &slice[0])
	fmt.Println(&nslice[1], &slice[1])
	fmt.Println(&nslice[2], &slice[2])

	//Todo

	fmt.Println("########")
	s2 := make([]int, 5)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", &s2[0])
	fmt.Println(len(s2), cap(s2))

	s2 = append(s2, 1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", &s2[0])
	fmt.Println(len(s2), cap(s2))

	s2 = append(s2, 1)
	fmt.Printf("%p\n", &s2[0])
	fmt.Printf("%p\n", s2)
	fmt.Println(len(s2), cap(s2))
}
