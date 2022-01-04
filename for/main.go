package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 5}
	k := []int{2, 3, 5}
	for _, v := range s {
		for _, j := range k {
			if v == j {
				fmt.Println("sas", v)
				break
			}
		}
	}
}
