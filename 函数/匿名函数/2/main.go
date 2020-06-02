package main

import "fmt"

//由 main 函数作为程序入口点启动
func main() {
	for i := 0; i < 3; i++ {
		//多次注册延迟调用，相反顺序执行
		defer func() {
			fmt.Println(i) //闭包引用局部变量
		}()

		fmt.Print(i)
		if i == 2 {
			fmt.Printf("\n")
		}
	}
}
