package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(Phone("121111"))
	fmt.Println(PhenCh("12345678900"))
}

func Phone(phone string) bool {
	reg, err := regexp.Compile(`[0-9]{0,14}$`)
	if err != nil {
		return false
	}
	return reg.Match([]byte(phone))
}

func PhenCh(phone string) bool {
	if len([]rune(phone)) != 11 {
		return false
	}
	reg, err := regexp.Compile(`/^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$/`)
	if err != nil {
		return false
	}
	return reg.Match([]byte(phone))
}
