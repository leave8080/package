package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	URLSchema            = "http://"
	URLHost              = "gfapi.mlogcn.com"
	URLGeopositionSearch = "/function/v001/poi"
	URLForairy1Hour      = "/air/v001/airhourly"
	URLWeatherReal       = "/air/v001/pastaqi"
	URLAirHour           = "/weather/v001/hour"
)

func GetAirFuture(ctx context.Context, areacode string) (rsp *AirQRsp, err error) {
	req := make(url.Values)
	req.Set("key", "W2jhBnzCp7gbb3JUs1gywrmFcLS6hbcZ")
	req.Set("hours", "25")
	req.Set("areacode", areacode)
	req.Set("output_type", "json")
	param := req.Encode()

	path := URLSchema + URLHost + URLForairy1Hour + "?" + param
	log.Println("path:", path)
	method := "GET"
	client := &http.Client{}
	Request, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}
	ress, err := client.Do(Request)
	if err != nil {
		return nil, err
	}
	defer ress.Body.Close()
	body, err := ioutil.ReadAll(ress.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Status == nil {
		log.Println("dsfsdafsdf")
	}
	return rsp, nil
}

func main() {
	air, err := GetAirFuture(context.Background(), "101090515")
	if err != nil {
		log.Println("dasd", err)
	}
	log.Println(air.Result.AqiHourlyFcst[0].Pm25, err)
}

type AirQRsp struct {
	Status *int `json:"status"`
	Result struct {
		AqiHourlyFcst []*AqiHourlyFcst `json:"aqi_hourly_fcst"`
	}
}

type AqiHourlyFcst struct {
	Aqi      int     `json:"aqi"`       //实时空气质量指数
	AqiLevel string  `json:"aqi_level"` //空气质量等级
	Pm10     int     `json:"pm10"`      //PM10浓度，单位: 微克/立方米
	Pm25     int     `json:"pm25"`      //PM2.5浓度，单位: 微克/立方米
	No2      int     `json:"no2"`       //二氧化氮浓度，单位: 微克/立方米
	So2      int     `json:"so2"`       //二氧化硫浓度，单位: 微克/立方米
	Co       float32 `json:"co"`        //一氧化碳浓度，单位: 毫克/立方米
	O3       int     `json:"o3"`        //臭氧浓度，单位: 微克/立方米
	//DataTime Time    `json:"data_time"` //预报时次
}

type Time time.Time
