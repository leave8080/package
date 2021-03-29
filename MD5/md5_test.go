package MD5

import (
	"fmt"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	fmt.Println(GeneratePassword("a111111", ""))
	fmt.Println(int(^uint(0) >> 1))
}
