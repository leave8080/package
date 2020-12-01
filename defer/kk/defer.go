package kk

import "fmt"

func Transaction() {
	//panicked := true

	if 1 < 2 {
		defer func() {
			fmt.Println("@@@@")
		}()
		fmt.Println("!!!!!")
	}
	fmt.Println("####")
}
