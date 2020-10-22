package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	for i := range s {
		s = append(s, i)
	}
	fmt.Println(s)
}

/*
请问如下程序是否能正常结束?

程序解释:main()函数中定义一个切片v，通过range遍历v，遍历过程中不断向v中添加新的元素。
参考答案:能够正常结束。循环内改变切片的长度，不影响循环次数，循环次效在循环开始前就已经确定了

*/
