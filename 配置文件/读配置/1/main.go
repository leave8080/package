package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.SetConfigFile("./config.yaml") // 指定配置文件路径
	// v.SetConfigName("config") // 配置文件名称(无扩展名)
	// v.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
	//viper.AddConfigPath(".")               // 还可以在工作目录中查找配置
	err := v.ReadInConfig() // 查找并读取配置文件
	if err != nil {         // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(v.Get("TimeStamp"), v.Get("Author"))
}
