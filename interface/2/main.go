package main

import "fmt"

type Pusher interface {
	Parse(mode int)
}

type SmsConfig struct {
}

func (J SmsConfig) Parse(mode int) {
	fmt.Println(mode)
}

type NoticeConfig struct {
}

func (Y NoticeConfig) Parse(mode int) {
	fmt.Println(mode)
}

func NewIRuleConfigParser(t string) Pusher {
	switch t {
	case "sms":
		return SmsConfig{}
	case "Notice":
		return NoticeConfig{}
	}
	return nil
}
func main() {
	code := "sms"
	ss := NewIRuleConfigParser(code)
	if ss != nil {
		ss.Parse(1)
	}

}
