package main

import "github.com/toolkits/sys"

func main() {
	//v, _ := mem.VirtualMemory()
	//res, err := cpu.Times(false)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res[0].Total(),res[0].User,res[0].System,res[0].Idle)
	//fmt.Printf("总内存: %v MB, 已使用:%v MB, 已使用百分比:%.f%%\n", v.Total>>20, v.Used>>20, ((res[0].Total()-res[0].Idle)/res[0].Total())*100)
	sys.CmdOut("/bin/sh", "-c", "reboot")
}
