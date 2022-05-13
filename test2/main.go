package main

import "fmt"

func main() {
	sMap := make(map[string]interface{})
	sMap["a"] = 1
	sMap["b"] = "c"
	sMap["c"] = 1.1
	sMap["d"] = true
	sMap["e"] = []int{1, 2, 3}
	for _, v := range sMap {
		s, err := InterfaceTostring(v)
		if err != nil {
			fmt.Println(err)

		}
		fmt.Printf("%T,%T\n", s, v)
		fmt.Println(s)
	}
}

func InterfaceTostring(v interface{}) (string, error) {

	switch v.(type) {
	case int:
		return fmt.Sprintf("%d", v), nil
	case string:
		return fmt.Sprintf("%s", v), nil
	case float64:
		return fmt.Sprintf("%f", v), nil
	case bool:
		return fmt.Sprintf("%t", v), nil
	case []int:
		return fmt.Sprintf("%v", v), nil
	default:
		return "", fmt.Errorf("%s", "type not support")
	}

}
