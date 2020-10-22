package main

import (
	"fmt"
	"math/rand"
	"sort"
	"unsafe"
)

type DownlownItem struct {
	Downlowntimes int
	Id            int
}
type DownloadCollection []*DownlownItem

func (this DownlownItem) String() string {
	return fmt.Sprintf("AppId:%d,DownloadTimes:%d", this.Id, this.Downlowntimes)
}

func (this DownloadCollection) Len() int {
	return len(this)
}
func (this DownloadCollection) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

// 根据app下载量降序排列
func (d DownloadCollection) Less(i int, j int) bool {
	return d[i].Downlowntimes > d[j].Downlowntimes
}

func main() {
	fmt.Println(unsafe.Sizeof(DownlownItem{}))
	a := make(DownloadCollection, 5)
	for i := 0; i < len(a); i++ {
		a[i] = &DownlownItem{i + 1, rand.Intn(1000)}
	}
	fmt.Println(a)
	sort.Sort(a)
	fmt.Println(a)
}
