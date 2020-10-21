package main

import (
	"fmt"
)

func main() {
	var arr = []int{19, 28, 17, 5, 13, 4, 6, 7, 9, 3, 10}
	//升序
	selectAscendingSort(arr)
}
func selectAscendingSort(ss []int) {
	l := len(ss)
	m := l - 1
	for i := 0; i < m; i++ {
		k := i
		for j := i + 1; j < l; j++ {
			if ss[k] > ss[j] {
				k = j
			}
		}
		if k != i {
			ss[i], ss[k] = ss[k], ss[i]
		}
	}
	fmt.Println(ss)
}
func test2(ss []int) {

	l := len(ss)

	for i := 0; i < l-1; i++ {
		k := i
		for j := i + 1; j < l; j++ {
			if ss[k] > ss[j] {
				k = j
			}
		}
		if k != i {
			ss[i], ss[k] = ss[k], ss[i]
		}
	}
}

/*
选择排序
选择排序(Selection sort) 是一种简单直观的排序算法。它的工作原理如下。首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。

选择排序的主要优点与数据移动有关。如果某个元素位于正确的最终位置上，则它不会被移动。选择排序每次交换一对元素，它们当中至少有一个将被移到其最终位置上，因此对n个元素的表进行排序总共进行至多n-1次交换。在所有的完全依靠交换去移动元素的排序方法中，选择排序属于非常好的一种。

算法原理
选择排序算法的原理如下：

首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置

再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。

重复第二步，直到所有元素均排序完毕。


*/
