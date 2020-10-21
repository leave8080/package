package main

import "fmt"

func main() {
	ss := []int{4, 6, 2, 2, 10, 5}
	quit(ss, 0, len(ss)-1)
	fmt.Println(ss)
}

func quit(ss []int, start, end int) {
	if start < end {
		i, j := start, end
		key := ss[(start+end)/2]
		for i <= j {
			for ss[i] < key {
				i++
			}
			for ss[j] > key {
				j--
			}
			if i <= j {
				ss[i], ss[j] = ss[j], ss[i]
				i++
				j--
			}
		}
		if start < j {
			quit(ss, start, j)
		}
		if end > i {
			quit(ss, i, end)
		}
	}

}
