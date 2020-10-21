package main

import "fmt"

/************************************
 *函数名：quicksort
 *作用：快速排序算法
 *参数：
 *返回值：无
 *模拟:
	begin:[]int{12, 85, 25, 16, 34, 23, 49, 95, 17, 61}
	-->[],12,[25 16 34 23 49 95 17 61]
	---->[23 16 17],25,[95 49 61 85]
	------>[17 16],23,[]
	-------->[16],17,[]
	---------->[],16,[]
	------>[34 49 61 85],95,[]
	-------->[],34,[61 85 95]
	---------->[49],61,[95]
	------------>[],49,[]
	------------>[85],95,[]
	-------------->[],85,[]
	last:[]int{12, 16, 17, 23, 25, 34, 49, 61, 85, 95}
 ************************************/

func quicksort(array []int, begin, end int, mark string) {
	var i, j int
	if begin < end {
		i = begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
		j = end       // array[end]是数组的最后一位

		for {
			if i >= j {
				break
			}
			if array[i] > array[begin] {
				array[i], array[j] = array[j], array[i]
				j = j - 1
			} else {
				i = i + 1
			}

		}

		/* 跳出while循环后，i = j。
		 * 此时数组被分割成两个部分  -->  array[begin+1] ~ array[i-1] < array[begin]
		 *                           -->  array[i+1] ~ array[end] > array[begin]
		 * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
		 * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件就退出！
		 */
		if array[i] >= array[begin] { // 这里必须要取等“>=”，否则数组元素由相同的值时，会出现错误！
			i = i - 1
		}

		array[begin], array[i] = array[i], array[begin]
		fmt.Printf("%s>%v,%d,%v\n", mark, array[begin:i], array[i], array[j:end])
		quicksort(array, begin, i, mark+"--")
		quicksort(array, j, end, mark+"--")
	}
}

func main() {

	nums := []int{12, 85, 25, 16, 34, 23, 49, 95, 17, 61}
	fmt.Printf("begin:%#v\n", nums)

	// 缩进
	mark := "--"
	quicksort(nums, 0, len(nums)-1, mark)
	fmt.Printf("last:%#v\n", nums)
}
