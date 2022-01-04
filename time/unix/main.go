package main

import (
	"fmt"
	"net/mail"
	"strconv"
	"time"
)

func main() {
	//date := 1616839155124
	//dateTime := time.Unix(int64(date/1000), int64(date%1000)*1000000)
	//
	//time.Unix(int64(date/1000), 0)
	//fmt.Println(time.Unix(int64(date/1000), 0))
	//
	//rer, err := msToTime(strconv.FormatInt(int64(date), 10))
	//if err != nil {
	//	log.Error("dad", rer)
	//	return
	//}
	//fmt.Println("dsada", rer)
	//fmt.Println(rer.Before(dateTime))

	fmt.Println("dsadad")

	fmt.Println(HourFormat(1639929600, 1640966400, 0))
	//fmt.Println(VerifyEmailFormat("154@qq.co.can"))
}

func msToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	tm := time.Unix(0, msInt*int64(time.Millisecond))

	fmt.Println(tm.Format("2006-02-01 15:04:05.000"))

	return tm, nil
}

func TimeFormate() {
	var s, v, i int64
	s, v = 1638288000, 1640880000
	fmt.Println(v - s)

	for i = s; i < v; i = RRR(i) {
		fmt.Println("dada", i)
		time.Unix(i, 0).Format("2006-01-02")
	}

}
func RRR(s int64) int64 {
	starTime := time.Unix(s, 0)
	starMonths := time.Date(starTime.Year(), starTime.Month(), starTime.Day(), 0, 0, 0, 0, starTime.Location())
	//fmt.Println("ss", starMonths.Format("2006/01/02 15:04:05"))
	retDay := starMonths.AddDate(0, 1, 0)
	//fmt.Println("eee", retDay.Format("2006/01/02 15:04:05"))
	return time.Date(retDay.Year(), retDay.Month(), retDay.Day(), 0, 0, 0, 0, retDay.Location()).Unix()
}

//正序开始时间
func monthAscStartTime(pagesize, pageNo int) int64 {
	now := time.Now()
	lastMonthFirstDay := now.AddDate(0, -(pagesize - 1), -now.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, now.Location())
	retDay := lastMonthStart.AddDate(0, pageNo-1, -lastMonthStart.Day()+1)

	return time.Date(retDay.Year(), retDay.Month(), retDay.Day(), 0, 0, 0, 0, lastMonthStart.Location()).Unix()
}

func VerifyEmailFormat(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func DayFormat(startTime, EndTime int64, orderType int) []string {

	list := make([]string, 0)
	for i := startTime; i < EndTime; i = DayStartAsc(i) {
		list = append(list, time.Unix(i, 0).Format("2006-01-02"))
	}

	return list
}

func DayStartAsc(s int64) int64 {
	starTime := time.Unix(s, 0)
	starMonths := time.Date(starTime.Year(), starTime.Month(), starTime.Day(), 0, 0, 0, 0, starTime.Location())
	//fmt.Println("ss", starMonths.Format("2006/01/02 15:04:05"))
	retDay := starMonths.AddDate(0, 0, 1)
	//fmt.Println("eee", retDay.Format("2006/01/02 15:04:05"))
	return time.Date(retDay.Year(), retDay.Month(), retDay.Day(), 0, 0, 0, 0, retDay.Location()).Unix()
}

func HourFormat(startTime, EndTime int64, orderType int) []string {

	list := make([]string, 0)
	for i := startTime; i < EndTime; i = HourStartAsc(i) {
		list = append(list, time.Unix(i, 0).Format("2006-01-02 15"))
	}
	//if orderType==model.OrderDesc{
	//倒叙
	//list=utils.SliceReplace(list)
	//}

	return list
}

func HourStartAsc(s int64) int64 {
	//starTime := time.Unix(s, 0)
	//starMonths := time.Date(starTime.Year(), starTime.Month(), starTime.Day(), 0, 0, 0, 0, starTime.Location())
	////fmt.Println("ss", starMonths.Format("2006/01/02 15:04:05"))
	//// starMonths.AddDate(0, 0, 0)
	//retDay := starMonths.AddDate(0, 0, 0).Add(1 * time.Hour)
	////fmt.Println("eee", retDay.Format("2006/01/02 15:04:05"))
	//return time.Date(retDay.Year(), retDay.Month(), retDay.Day(), 0, 0, 0, 0, retDay.Location()).Unix()
	return s + 3600
}
