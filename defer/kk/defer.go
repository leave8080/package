package kk

import (
	"fmt"
	"log"
)

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

func DeferFunc() {
	defer func() {
		fmt.Println("defer")
	}()
	log.Fatal("fatal")
}
