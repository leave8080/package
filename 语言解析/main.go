package main

import (
	"fmt"
	"golang.org/x/text/language"
)

func main() {

	lang, err := language.Parse("zh-Hans")
	if err != nil {
		lang = language.MustParse("zh")
		err = nil
	}
	//fmt.Println(lang.Parent(), language.Und)
	if lang.Parent() != language.Und {
		lang = lang.Parent()
	}
	fmt.Println(lang.String())
}
