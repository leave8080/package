package main

import (
	"github.com/leave8080/package/test/rr"

	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Environ())
	rr.EE()
}
