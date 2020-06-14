package main

import "fmt"

func main1() {
	testList := []uint16{1, 2, 3, 4}
	fmt.Printf("%p\n", &testList)
	myTest := testList //1.mytest 和 testList指向同一段
	fmt.Println(len(myTest), cap(myTest))
	testList[0] = 10
	fmt.Printf("%p\n", &myTest)
	myTest = append(myTest, 9) //分离指向,所以在设计的时候最好给一个cap，不需要频繁的开辟内存
	fmt.Printf("%p\n", &myTest)
	fmt.Println(len(myTest), cap(myTest))
	//myTest[0] = 9
	//fmt.Println(testList)
	//fmt.Println(myTest)
}
func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := a[1:6]
	fmt.Printf("%p\n", &s)
	s = append(s, 1, 2, 3)
	fmt.Printf("%p\n", &s)
	for i := 0; i < 10; i++ {
		fmt.Println(cap(s))
		s = append(s, i)
	}
}
