package main

import (
	"fmt"
	"sort"
)

func main() {
	str := make(map[int]int, 4)
	str[1] = 10
	str[2] = 3
	str[3] = 5
	str[4] = 33
	for i, v := range str {
		fmt.Println(i, v)
	}
	fmt.Println()
	s := make([]int, 0)
	for i, _ := range str {
		s = append(s, i)
	}
	sort.Ints(s)
	for _, v := range s {
		fmt.Println(v, str[v])
	}

}
