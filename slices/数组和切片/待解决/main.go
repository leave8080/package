package main

import "fmt"

func main1() {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		fmt.Println("4:", a)
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Print(a)
		}
		fmt.Println("2:", v, &a[1])
		a[k] = 100 + v
		fmt.Println("3:", a)
	}
	fmt.Print(a)
}

/*
	100 200 3
	101
	102
	103
*/

func main2() {
	a := []int{1, 2, 3}
	fmt.Println(&a[1])
	for k, v := range a {

		fmt.Println("1:", v, &a[1])
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Print(a)
		}
		fmt.Println("2:", v, &a[1])
		a[k] = 100 + v
		fmt.Println("3:", a)
	}
	fmt.Print(a)
}
func main() {
	main1()
	//main2()
}

/*
	100,200,3

*/
