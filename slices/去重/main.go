package main

import "fmt"

func main() {
	ss := []int{1, 4, 56, 7, 4, 7, 1}
	if len(ss) < 1024 {

		RemoveRepByLoop(ss)
	} else {
		RemoveRepByMap(ss)
	}

}
func RemoveRepByLoop(ss []int) {
	result := []int{} // 存放结果

	for i := range ss {
		flag := true
		fmt.Println("1")
		for j := range result {
			fmt.Println("2")
			fmt.Println(result)
			if ss[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			fmt.Println("3")
			result = append(result, ss[i])
		}
	}
	//fmt.Println(result)
}

func RemoveRepByMap(ss []int) {
	Maps := make(map[int]int)
	for _, v := range ss {
		Maps[v] = 0
	}
	ss = ss[0:0]
	for i := range Maps {
		ss = append(ss, i)
	}

	fmt.Println(ss)
}
