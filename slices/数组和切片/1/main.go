package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a)
			//fmt.Println("!!!!!!!!")
		}

		a[k] = 100 + v
		//fmt.Println(a, k)
	}
	//fmt.Println()
	fmt.Print(a)
}

func main2() {
	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a)
			fmt.Println("!!!!!!!!")
		}

		a[k] = 100 + v
		fmt.Println(a, k)
	}
	//fmt.Println()
	fmt.Print(a)
}
