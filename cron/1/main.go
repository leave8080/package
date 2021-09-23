package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func task() {
	fmt.Println("I am running task.")
}

func task2() {
	fmt.Println("I am running task2.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	Run()
	for {
		fmt.Println(time.Now())
		time.Sleep(time.Second * 20)
	}
}

func Run() {

	//s2 := gocron.NewScheduler(time.FixedZone("UTC-8", +8*60*60))
	//l, _ := time.LoadLocation("Asia/Shanghai")
	m := time.FixedZone("UTC", 8*3600)
	s2 := gocron.NewScheduler(m)

	_, err := s2.Every(1).Day().At("09:39").Do(task)
	if err != nil {
		log.Fatalf("error creating job: %v", err)
	}

	// executes the scheduler and blocks current thread
	//s2.StartBlocking()
	s2.StartAsync()
	// this line is never reached
}
