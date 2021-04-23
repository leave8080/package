package main

import (
	"fmt"
	"time"
)

const MAXSLEEP = 128

func main() {
	for numsec := 1; numsec <= MAXSLEEP; numsec <<= 1 {
		// TODO
		fmt.Println(numsec)
		if numsec <= MAXSLEEP/2 {
			time.Sleep(time.Second * time.Duration(numsec))
			fmt.Println("slepp time(s):", numsec)
		}
	}
}
