package main

import (
	"fmt"
	"strconv"
)

func main() {
	dd := make(map[int][]int)
	ddd := make(map[int][]int, 0)
	for i := 0; i < 3; i++ {
		dd[i] = append(dd[i], i)
		ddd[i] = append(ddd[i], i)
	}

	fmt.Println(dd)
	fmt.Println(ddd)

	var rr interface{}
	rr = 10
	fmt.Println(rr)
	fmt.Printf("%T\n", rr)
	if e, ok := rr.(int); ok {
		fmt.Println(e)
	}

	ssss, err := strconv.ParseFloat("1.1", 64)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(ssss)
	}
	fmt.Println("dadadaada")
	sd := make(map[string]interface{})
	sd["a"] = 1
	sd["b"] = "c"
	sd["c"] = 1.1
	sd["d"] = true
	sd["e"] = []int{1, 2, 3}

	//for i, v := range sd {
	//	if _, ok := v.(string); ok {
	//		fmt.Println(v)
	//	} else {
	//		fmt.Println(">>>>>:", i)
	//	}
	//}
	for i, v := range sd {
		s, err := InterfaceTostring(v)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("%T\n", v)
			continue
		}
		fmt.Printf("%T,%T\n", s, v)
		fmt.Println(s, i)
	}
}

func InterfaceTostring(sValue interface{}) (string, error) {
	switch sValue.(type) {
	case string:
		return sValue.(string), nil
	case int:
		return strconv.Itoa(sValue.(int)), nil
	case int32:
		return strconv.Itoa(int(sValue.(int32))), nil
	case int64:
		return strconv.Itoa(int(sValue.(int64))), nil
	case float64:
		return strconv.FormatFloat(sValue.(float64), 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(sValue.(float32)), 'f', -1, 32), nil
	case bool:
		return strconv.FormatBool(sValue.(bool)), nil
	default:
		return "", fmt.Errorf("type error no suppurt %T", sValue)
	}
}
