package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func main() {
	ExampleJob_NextRun()
}

func ExampleJob_NextRun() {
	s := gocron.NewScheduler(time.UTC)
	job, _ := s.Every(1).Second().Do(Task)
	go func() {
		for {
			fmt.Println("Next run", job.RunCount())
			time.Sleep(time.Second)
		}
	}()
	s.StartAsync()
	select {}
}
func Task() {
	//fmt.Println("time:", time.Now())
}
