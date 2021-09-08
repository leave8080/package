package main

import (
	"fmt"
	"gofer/pkg/log"
	"golang.org/x/text/language"
)

func main() {
	lang, err := language.Parse("zh-CN")
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println(lang.String())

	var (
		values []interface{}
	)
	values = append(values, 1, 2, "dasd")
	fmt.Println(values)
}
