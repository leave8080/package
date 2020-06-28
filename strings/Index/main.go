package main

import (
	"fmt"
	"strings"
)

func main() {
	filename := "abc.beta1.window"
	ss := StringValue(filename, "beta")
	fmt.Println(ss)
}

func StringValue(origin string, param string) string {
	i := strings.Index(origin, param)
	if i != -1 {
		j := i + len(param)
		k := strings.Index(origin[j:], ".window")
		return param + strings.TrimSpace(origin[j:j+k]) + "."
	}
	return ""
}
