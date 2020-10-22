package main

import "fmt"

func main() {

	ss := []int{1, 2, 3, 4, 5}
	ss = append(ss[0:1], ss[2:]...)
	fmt.Println(ss)
}
