package main

import "fmt"

func main() {
	var name = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ss := name[5:6]
	fmt.Println(&ss[0])
	ss = append(ss, 1, 2, 3, 4)
	fmt.Println(cap(ss), &ss[0])
	fmt.Println("扩容=========")
	ss = append(ss, 5)
	fmt.Println(cap(ss), &ss[0])

}

//0xc000018118
//5 0xc000018118
//扩容=========
//10 0xc000018140

//⚠️ cap
