package main

import "fmt"

type Push interface {
	PushCOde()
	Fs
}

type Fs interface {
}
type SMs struct {
}
type Mob struct {
}

func (this *SMs) PushCOde() {
	fmt.Println("sms")
}

func (this *Mob) PushCOde() {
	fmt.Println("mob")

}
func main() {
	s := &SMs{}
	m := &Mob{}
	//var key int
	key := 1
	switch key {
	case 1:
		s.PushCOde()
	case 2:
		m.PushCOde()
	}
}
