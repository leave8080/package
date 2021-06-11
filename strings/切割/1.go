package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "weqqeqeqeqe"
	sss := strings.Split(ss, "_")
	fmt.Println(sss[0])

	sl := make([][]int, 2)
	sl[0] = []int{1, 2}
	sl[1] = []int{3, 4}
	fmt.Println(sl)
}
