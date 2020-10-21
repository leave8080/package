package main

import (
	"fmt"
	"sort"
)

func sorte(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}
func sort2(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
func sort3(s []int) []int {
	a := []int{4, 3, 2, 1, 5, 9, 8, 7, 6}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println("After reversed: ", a)
	return a
}

func sort4(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
func main() {
	ss := []int{0, 1, 6, 3, 6, 7, 2}

	ss = sort4(ss)
	fmt.Println(ss)
}
