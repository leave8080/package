package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Scope string
}

type DeviceInfo struct {
	ProductKey string `json:"product_key"`
	DeviceName string `json:"device_name"`
}

func main() {
	//	s := &Data{
	//		Scope: `{
	//    "product_key":"a1Rv30rqHka",
	//    "device_name":"2bd5B0m1xDXLcRQjj17x"
	//}
	//`,
	//	}
	ss := &DeviceInfo{
		ProductKey: "dsad",
		DeviceName: "we",
	}

	by, err := json.Marshal(ss)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(s.Scope)
	info := &DeviceInfo{}
	if err := json.Unmarshal(by, info); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info.ProductKey, info.DeviceName)
}
