package age

import (
	"gofer/pkg/log"
	"strings"
	"time"
)

/*
根据出生年月，计算年龄
*/
func GetYearDiffer(start_time, end_time string) int {
	var Age int64
	var lasr int64
	var pslTime string
	if strings.Index(start_time, ".") != -1 {
		pslTime = "2006.01.02"
	} else if strings.Index(start_time, "-") != -1 {
		pslTime = "2006-01-02"
	} else {
		pslTime = "2006/01/02"
	}
	t1, err := time.ParseInLocation(pslTime, start_time, time.Local)
	t2, err := time.ParseInLocation(pslTime, end_time, time.Local)
	log.Info("t1:", t1)
	log.Info("t2:", t2)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		log.Info("diff:", diff)
		Age = diff / (3600 * 365 * 24)
		lasr = diff % (3600 * 365 * 24)

		log.Info("lasr:", lasr)
		log.Info("age:", Age)
		log.Info("moun:", lasr/(3600*30*24))
		return int(Age)
	} else {
		return int(Age)
	}
}

/*

func GwtAge(start_time, end_time string) string {
	var pslTime string
	if strings.Index(start_time, ".") != -1 {
		pslTime = "2006.01.02"
	} else if strings.Index(start_time, "-") != -1 {
		pslTime = "2006-01-02"
	} else {
		pslTime = "2006/01/02"
	}

	return ""
}





*/
