package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func main() {
	fmt.Println(time.Now())
	s := gocron.NewScheduler(time.UTC)
	//s.Cron("*/1 * * * *").Do(Task1) // every minute
	_, err := s.Cron("0 * * * *").Do(Task)
	if err != nil {
		fmt.Println("11", err)
	}
	//s.Cron("* * * * * ").Do(Task3)
	//_, err = s.Every("1h").Hours().Do(Task2)
	//if err != nil {
	//	fmt.Println("22", err)
	//}
	//s.Every(1).Second().Do(Task)
	//s.Every(5).Second().Do(Task2)
	s.StartAsync() //异步
	// starts the scheduler and blocks current execution path
	//s.StartBlocking() //同步

	for {
		//fmt.Println("1")
		time.Sleep(time.Second * 10)
	}
}

func Task1() {
	fmt.Println("task1", time.Now())
}
func Task() {
	fmt.Println("time:", time.Now())
}
func Task2() {
	fmt.Println("TASK2", time.Now())
}

func Task3() {
	fmt.Println("task3", time.Now())
}
