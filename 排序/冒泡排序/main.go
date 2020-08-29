package main

import "fmt"

func main() {
	ss := []int{1, 7, 3, 6, 2, 7, 8, 33, 34, 23}
	fmt.Println("before:", ss)
	sort1(ss)
	fmt.Println("after:", ss)
}

//优化1 适用于 绝大数已经排好序
func sort1(ss []int) {
	for i := 0; i < len(ss)-1; i++ {
		Issort := true
		for j := 0; j < len(ss)-i-1; j++ {
			if ss[j] > ss[j+1] {
				ss[j], ss[j+1] = ss[j+1], ss[j]
				Issort = false
			}

		}
		if Issort {
			break
		}
	}

}

func sort2(ss []int) {
	for i := 0; i < len(ss)-1; i++ {
		Issort := true
		for j := 0; j < len(ss)-i-1; j++ {
			if ss[j] > ss[j+1] {
				ss[j], ss[j+1] = ss[j+1], ss[j]
				Issort = false
			}

		}
		if Issort {
			break
		}
	}

}

/*
算法原理
冒泡排序算法的原理如下：

比较相邻的元素。如果第一个比第二个大，就交换他们两个。
对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数。
针对所有的元素重复以上的步骤，除了最后一个。
持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
*/
