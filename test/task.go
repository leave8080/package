package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	const num = 2
	s := 1609557368 //2021-01-02 11:16:08
	e := 1619838968 //2021-05-02 11:16:08
	t := (e - s) / (60 * 60 * 24 * 30 * num)
	//t := (e - s) / (60 * 60 * 24 * 30)
	fmt.Println((e - s) / (60 * 60 * 24 * 30))            //月
	fmt.Println((e - s - t*30*24*3600*num) / (24 * 3600)) //日
	local := s + t*24*3600*30*num
	fmt.Println(time.Unix(int64(local), 0).Format("2006-01-02 15:04:05"))

	rsp, err := http.Get("https://dev-app.api.cloud.mxchip.com/app/v1/ping")
	if err != nil {
		fmt.Println("get err", err)
		//fcLogger.Error("get baidu err",err)
		//return event,err
	}
	fmt.Println("rsp", rsp)

}
