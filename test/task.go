package main

import "fmt"

func main() {
	messNum := 128
	var messIds []int
	for i := 0; i < 12; i++ {
		//if i == 4 {
		//	continue
		//}
		if messNum&1 == 1 {
			messIds = append(messIds, i)
		}
		messNum >>= 1
	}
	fmt.Println(messIds)
	messIds = []int{1, 2, 3, 4, 5, 8}
	messIds = DeleteSlice(messIds)
	fmt.Println(messIds)
}

func DeleteSlice(a []int) []int {
	for i := 0; i < len(a); i++ {
		if a[i] == 8 {
			a = append(a[:i], a[i+1:]...)
			i--
		}
	}
	return a
}
func DeleteSlice2(a []int) []int {
	j := 0
	for _, val := range a {
		if val == 8 {
			a[j] = val
			j++
		}
	}
	return a[:j]
}
