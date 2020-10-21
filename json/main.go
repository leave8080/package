package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

var ExTimeUnmarshalTimeFormat = "2006-01-02 15:04:05"
var ExTimeMarshalTimeFormat = "2006-01-02 15:04:05"

type ExTime struct {
	time.Time
}

func (t ExTime) UnmarshalTimeFormat() string {
	return ExTimeUnmarshalTimeFormat
}

func (t ExTime) MarshalTimeFormat() string {
	return ExTimeMarshalTimeFormat
}

type Event struct {
	EventId int64  `json:"id"`
	Stamp   ExTime `json:"stamp"`
}

func (t *ExTime) UnmarshalJSON(b []byte) error {
	b = bytes.Trim(b, "\"")
	ext, err := time.Parse(t.UnmarshalTimeFormat(), string(b))
	if err != nil {
		// do something
	}
	*t = ExTime{ext}
	return nil
}

func (t ExTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t.Time).Format(t.MarshalTimeFormat()))
	return []byte(stamp), nil
}

func main() {
	e1 := `{"id":5,"stamp":"2019-01-02 15:04:05"}`
	p := new(Event)
	_ = json.Unmarshal([]byte(e1), p)
	fmt.Println(p.Stamp.Format(ExTimeMarshalTimeFormat)) // output : 2006-01-02 15:04:05 +0000 UTC

	e2 := `{"id":5,"stamp":"Mon Jan 2 15:04:05 200=16"}`
	ExTimeUnmarshalTimeFormat = "Mon Jan _2 15:04:05 2006"
	p2 := new(Event)
	_ = json.Unmarshal([]byte(e2), p2)
	fmt.Println(p2.Stamp) // output : 2006-01-02 15:04:05 +0000 UTC

	ej1, _ := json.Marshal(p)
	fmt.Println(string(ej1)) // output : {"id":5,"stamp":"2006-01-02 15:04:05"}

	ExTimeMarshalTimeFormat = "Mon Jan _2 15:04:05 2006"
	ej2, _ := json.Marshal(p)
	fmt.Println(string(ej2)) // output : {"id":5,"stamp":"Mon Jan  2 15:04:05 2006"}
}
