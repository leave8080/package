package main

import "fmt"

func main() {
	orderlen := 5
	order := make([]uint16, 2*orderlen)

	pollroder := order[:orderlen:orderlen]
	lockorder := order[orderlen:][:orderlen:orderlen]
	fmt.Println(len(pollroder))
	fmt.Println(cap(pollroder))
	fmt.Println(len(lockorder))
	fmt.Println(cap(lockorder))
	ss := []int{1, 2, 3, 4}
	sss := ss[2:][:2:2]
	fmt.Println(sss)
}
