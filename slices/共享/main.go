package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a[:2]
	a[0] = 10
	//此时a和b共享底层数组
	fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
	fmt.Println(b, "b cap:", cap(b), "b len:", len(b))
	fmt.Println("-----------")
	b = append(b, 999)
	//虽然b append了1,但是没有超出cap，所以未进行内存重新分配
	//等同于b[2]=999，因此a[2]一并被修改
	fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
	fmt.Println(b, "b cap:", cap(b), "b len:", len(b))
	fmt.Println("-----------")
	a[2] = 555 //同上，未重新分配，所以，a[2] b[2]都被修改
	fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
	fmt.Println(b, "b cap:", cap(b), "b len:", len(b))
	fmt.Println("-----------")
	b = append(b, 777) //超出了cap，这时候b进行重新分配,b[3]=777,cap(b)=6
	a[2] = 666         //这时候a和b不再共享
	fmt.Println(a, "a cap:", cap(a), "a len:", len(a))
	fmt.Println(b, "b cap:", cap(b), "b len:", len(b))
}
