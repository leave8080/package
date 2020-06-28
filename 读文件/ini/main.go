package main

import (
	"bufio"
	"fmt"
	"github.com/go-ini/ini"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/Users/leave/work/Gowork/src/github.com/leave8080/package/读文件/ini/Version.ini")
	if err != nil {
		fmt.Println("open err:", err)
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	keymap := make(map[string]string)
	for {
		l, err := buf.ReadString('\n')
		line := strings.TrimSpace(l)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read err", err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
		default:
			i := strings.IndexAny(line, "=")
			value1 := strings.TrimSpace(line[0:i])
			value := strings.TrimSpace(line[i+1 : len(line)])
			fmt.Println(value1, "=", value)
			keymap[value1] = value

		}
	}
	fmt.Println(keymap)
	for i, _ := range keymap {
		if i == "version" {
			keymap[i] = "V3.3.3.20200303"
		}
		tracefile(i + "=" + keymap[i])
	}
	fmt.Println(keymap)

}
func tracefile(str_content string) {
	fd, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd_content := strings.Join([]string{str_content, "\n"}, "")
	buf := []byte(fd_content)
	fd.Write(buf)
	fd.Close()
}

func main2() {
	cfg, err := ini.Load("/Users/leave/work/Gowork/src/github.com/leave8080/package/读文件/ini/Version.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	// 典型读取操作，默认分区可以使用空字符串表示
	fmt.Println("App Name:", cfg.Section("").Key("app_name").String())
	fmt.Println("Log Level:", cfg.Section("").Key("log_level").String())
}
