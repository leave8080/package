package main

import "fmt"

func main() {
	ss := []int{77, 41, 2, 3, 66, 33, 23, 7}
	ss = sortt(ss)
	fmt.Println(ss)
}
func sortt(ss []int) []int {
	for i := 0; i < len(ss)-1; i++ {
		key := i
		for j := i + 1; j < len(ss); j++ {
			if ss[j] < ss[key] {
				key = j
			}
		}
		fmt.Println(key, i)
		if key != i {
			ss[key], ss[i] = ss[i], ss[key]
		}
		fmt.Println(ss)
	}
	return ss
}
