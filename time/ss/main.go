package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix()*1000 + 3600000)
	fmt.Println(time.Now().Format("2006-01-02 15:00:00"))
}
