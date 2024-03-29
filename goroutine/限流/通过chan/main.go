package main

import (
	"fmt"
	"time"
)

func Run(task_id, sleeptime, timeout int, ch chan string) {
	ch_run := make(chan string)
	go run(task_id, sleeptime, ch_run)
	select {
	case re := <-ch_run:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d , timeout", task_id)
		ch <- re
	}
}

func run(task_id, sleeptime int, ch chan string) {

	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
	return
}

func main() {
	//input := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	input := []int{1, 1, 1, 1, 1, 1}
	timeout := 2
	chLimit := make(chan bool, 1)
	chs := make([]chan string, len(input))
	limitFunc := func(chLimit chan bool, ch chan string, task_id, sleeptime, timeout int) {
		Run(task_id, sleeptime, timeout, ch)
		<-chLimit
	}
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		chs[i] = make(chan string, 1)
		chLimit <- true
		go limitFunc(chLimit, chs[i], i, sleeptime, timeout)
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
}
