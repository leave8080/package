package replaces

import (
	"fmt"
	"strings"
)

func Replace() {
	obj := strings.Replace("/www", "api", "", 1)
	fmt.Println(obj)
}
