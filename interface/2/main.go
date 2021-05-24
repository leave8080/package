package main

import "fmt"

// IRuleConfigParser IRuleConfigParser
type Pusher interface {
	Parse(mode int)
}

// jsonRuleConfigParser jsonRuleConfigParser
type SmsConfig struct {
}

// Parse Parse
func (J SmsConfig) Parse(mode int) {
	fmt.Println(mode)
}

// yamlRuleConfigParser yamlRuleConfigParser
type NoticeConfig struct {
}

// Parse Parse
func (Y NoticeConfig) Parse(mode int) {
	fmt.Println(mode)
}

// NewIRuleConfigParser NewIRuleConfigParser
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
