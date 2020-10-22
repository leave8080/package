package main

import "fmt"

func main() {
	s1 := []int{1, 2}
	s2 := []int{3, 4, 5}
	fmt.Println(&s1[0], &s2[0])
	copy(s1, s2)
	fmt.Println(s1, cap(s1))
	fmt.Println(&s1[0], &s2[0])
	//[3 4] 2
	fmt.Println("@@@")
	s3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s4 := s3[1:4]
	fmt.Println(s4)
	s5 := s3[2:5]
	fmt.Println(s5)
	copy(s4, s5)

	fmt.Println(s4)
	fmt.Println(s5)

	fmt.Println("##")
	q1 := []int{1, 2, 3}
	q2 := []int{4, 5, 6, 7}
	copy(q2, q1)
	fmt.Println(q1)
	fmt.Println(q2)
	test()
}
func test() error {

	return fmt.Errorf("err")
}
