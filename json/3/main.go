package main

import (
	"encoding/json"
	"gofer/pkg/log"
)

func SpecJson(params string) (*TrggerParam, error) {
	spec := new(TrggerParam)
	if err := json.Unmarshal([]byte(params), &spec); err != nil {
		log.Errorf(" json.Unmarshal(%+v); err:%+v", params, err)
		return nil, err
	}
	return spec, nil
}
func Json2(pa string) (interface{}, error) {
	var i interface{}
	if err := json.Unmarshal([]byte(pa), &i); err != nil {
		return nil, err
	}
	return i, nil
}
func main() {
	list, err := Json2(`{"value": 1, "identifier": "Gas_Immersion", "compare_type": "=="}`)
	if err != nil {
		log.Errorf("err:%+v", err)
		return
	}
	log.Infof("%+v", list)
	log.Info(list.(map[string]interface{})["value"])
}

type TrggerParam struct {
	Identifier  string `json:"identifier" bson:"identifier"`
	CompareType string `json:"compare_type" bson:"compare_type"`
	Value       string `json:"value" bson:"value"`
}
