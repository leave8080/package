package age

import (
	"fmt"
	"testing"
)

func TestGetYearDiffer(t *testing.T) {
	fmt.Println(GetYearDiffer("2020.07.20", "2021.07.20"))
}
