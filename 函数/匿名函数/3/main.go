package main

import "fmt"

func A() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main1() {
	f1 := A()
	fmt.Println(f1())
	fmt.Println(f1())
}

func main() {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v) //每次将变量 v 的拷贝传进函数
	}
	select {}
}
