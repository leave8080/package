package lru

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}

func TestConstructor(t *testing.T) {
	Constructor(1)
}
