package main

import "fmt"

//golang为uint32、uint64、string提供了fast access，使用这些类型作为key可以提高map访问速度。[runtime/hashmap_fast.go]
//判断map 键值存在
func main1() {
	yinzhengjie := make(map[int]string)
	letter := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for k, v := range letter {
		yinzhengjie[k] = v
	}
	for k, v := range yinzhengjie {
		fmt.Println(k, v) //注意，字典是无序的哟！
	}
	//fmt.Printf("字典中的值为：【%v】\n", yinzhengjie)
	fmt.Println(yinzhengjie)
	if v, ok := yinzhengjie[1]; ok {
		fmt.Println("存在key=", v)
	} else {
		fmt.Println("没有找到key=", v)
	}

	v, ok := yinzhengjie[1]
	if ok {
		fmt.Println("再一次确认，已经存在key=", v)
	} else {
		fmt.Println("再一次确认，没有找到key=", v)
	}
}

func main2() {
	x := make(map[int]int)
	for i := 0; i < 30; i++ {
		x[i] = i
	}
	for k, v := range x {
		fmt.Println(k, v)
	}

}

func main() {
	var m map[int64]int64
	m = make(map[int64]int64, 1)
	fmt.Printf("m 原始的地址是：%p\n", m)
	change(m)
	fmt.Printf("m 改变后地址是：%p\n", m)
	fmt.Println("m 长度是", len(m))
	fmt.Println("m 参数是", m)
}

func change(m map[int64]int64) {
	fmt.Printf("m 函数开始时地址是：%p\n", m)
	var max = 5
	for i := 0; i < max; i++ {
		m[int64(i)] = 2
	}
	fmt.Printf("m 在函数返回前地址是：%p\n", m)
}
