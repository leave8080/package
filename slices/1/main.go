package main

import "fmt"

//https://blog.csdn.net/u013276277/article/details/103653066?utm_medium=distribute.pc_relevant_right.none-task-blog-BlogCommendFromBaidu-7.nonecase&depth_1-utm_source=distribute.pc_relevant_right.none-task-blog-BlogCommendFromBaidu-7.nonecase
func main() {
	//Subslice()
	subSlice2()
}
func Subslice() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ss := s[3:6]
	sss := s
	fmt.Printf("len ss : %d\n", len(ss))
	fmt.Printf("Cap ss : %d\n", cap(ss))
	ss = append(ss, 4)
	sss[1] = 444
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", &ss)
	fmt.Printf("%p\n", &sss)
	fmt.Println(sss)
	fmt.Println(ss)
	fmt.Println(s)
	fmt.Printf("len ss : %d\n", len(ss))
	fmt.Printf("Cap ss : %d\n", cap(ss))
	for _, v := range ss { //4,5,6,4
		v += 10 //14,15,16,14
	}
	fmt.Println(ss)
	for i := range ss {
		ss[i] += 10
	}

	fmt.Println(s)
}
func subSlice2() {
	s := []int{1, 2, 3}
	ss := s[1:]
	fmt.Println(len(ss), cap(ss))
	ss = append(ss, 4)
	fmt.Println(len(ss), cap(ss))
	for _, v := range ss {
		v += 10
	}
	fmt.Println(ss)
	for i := range ss {
		ss[i] += 10
	}
	fmt.Println(s)
	fmt.Println(ss)
}
