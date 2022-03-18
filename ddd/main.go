package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(VersionValidate("1.0.0"))
}
func VersionValidate(version string) bool {
	re := `^(\d+)\.(\d+)\.(\d+)$`
	match, err := regexp.MatchString(re, version)
	if err != nil {
		return false
	}
	return match
}
