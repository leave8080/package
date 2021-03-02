package Encoder

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	// 如果Child字段为nil 编码JSON时可忽略
	Child *Person `json:"child,omitempty"`
}

func Encoder() {
	person := Person{
		Name: "John",
		Age:  40,
		Child: &Person{
			Name: "Jack1",
			//Age:  20,
		},
	}

	// File类型实现了io.Writer接口
	file, _ := os.Create("person.json")

	// 根据io.Writer创建Encoder 然后调用Encode()方法将对象编码成JSON
	json.NewEncoder(file).Encode(&person)
}

func Decoder() {
	var person Person

	// File类型也实现了io.Reader接口
	file, _ := os.Open("person.json")

	// 根据io.Reader创建Decoder 然后调用Decode()方法将JSON解析成对象
	json.NewDecoder(file).Decode(&person)

	fmt.Println(person)
	fmt.Println(*person.Child)
}
