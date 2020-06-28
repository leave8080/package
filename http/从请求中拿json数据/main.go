package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonhandler(w http.ResponseWriter, r *http.Request) {
	cvalue := r.Header.Get("Content-Type")
	fmt.Println(cvalue)          //请求体的内容类型
	fmt.Println(r.ContentLength) //请求体的长度
	data := make([]byte, r.ContentLength)
	r.Body.Read(data)
	user := new(User)
	json.Unmarshal(data, user)
	fmt.Println(user)
}
func main() {
	http.HandleFunc("/", jsonhandler)
	http.ListenAndServe(":8080", nil)
}
