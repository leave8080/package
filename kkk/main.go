package main

import (
	"encoding/json"
	"gofer/pkg/log"
)

func main() {
	spec, err := SpecJson2(`{"0": "自动", "1": "制冷", "2": "制热"}`)
	if err != nil {
		log.Errorf("SpecJson(%+v); err:%+v", `{"0": "自动", "1": "制冷", "2": "制热"}`, err)
		return
	}
	log.Infof("spec:%+v", spec)
	//var ss map[string]interface{}
	ss := make(map[string]string, 0)
	ss["0"] = spec["0"]
	ss["1"] = spec["1"]
	//ss["2"] = spec["2"]
	log.Info(ss)
}
func SpecJson(propSpecs string) (spec interface{}, err error) {
	spec = new(interface{})
	if err := json.Unmarshal([]byte(propSpecs), &spec); err != nil {
		log.Errorf(" json.Unmarshal(%+v); err:%+v", propSpecs, err)
		return nil, err
	}
	return spec, nil
}

func SpecJson2(propSpecs string) (spec map[string]string, err error) {
	spec = make(map[string]string, 0)
	if err := json.Unmarshal([]byte(propSpecs), &spec); err != nil {
		log.Errorf(" json.Unmarshal(%+v); err:%+v", propSpecs, err)
		return nil, err
	}
	return spec, nil
}
