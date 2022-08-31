package main

import (
	"fmt"
	"strconv"
)

type CentralControlData struct {
	IotId   *[]string `json:"iotid,omitempty"`
	SceneId *[]int    `json:"sceneId,omitempty"`
	GroupId *[]int    `json:"groupId,omitempty"`
}

func main() {
	ss := &CentralControlData{
		IotId:   nil,
		SceneId: &[]int{1, 2, 3},
		GroupId: &[]int{},
	}
	fmt.Println(ss.IotId, ss.SceneId, ss.GroupId)
	//meshAddress, _ := strconv.Atoi("77")
	//fmt.Printf("%d", meshAddress)
	fmt.Println(Hex2Dec("77"))
}
func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	return int(n)
}
