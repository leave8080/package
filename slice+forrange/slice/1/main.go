package main

import "fmt"

func main() {
	nums := [3]int{}
	nums[0] = 1

	fmt.Printf("nums: %v , len: %d, cap: %d\n", nums, len(nums), cap(nums))

	dnums := nums[0:2]
	dnums[0] = 5

	fmt.Printf("nums: %v ,len: %d, cap: %d\n", nums, len(nums), cap(nums))
	fmt.Printf("dnums: %v, len: %d, cap: %d\n", dnums, len(dnums), cap(dnums))
}
