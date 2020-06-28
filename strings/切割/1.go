package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "V1.0.0.20200202"
	sss := strings.Split(ss, "-beta")
	fmt.Println(sss[0])
}
