package main

import "fmt"

func adder() func() {
	sum := 0
	f := func() {
		sum += 1
		fmt.Printf("闭包sum：%v", sum, "%p", &sum)
		fmt.Println()
	}
	fmt.Println("ad")
	return f
}

func main() {
	pos := adder()
	pos() //闭包sum：1
	pos() //闭包sum：2
}
