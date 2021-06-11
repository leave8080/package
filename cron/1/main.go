package main

import (
	"fmt"
	"log"

	//"log"
	"time"

	"github.com/go-co-op/gocron"
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
	// defines a new scheduler that schedules and runs jobs
	//s1 := gocron.NewScheduler(time.UTC)

	//s1.Every(3).Seconds().Do(task)

	// scheduler starts running jobs and current thread continues to execute
	//s1.StartAsync()

	// Do jobs without params
	s2 := gocron.NewScheduler(time.UTC)
	log.Println(time.UTC)
	//s2.Every(1).Second().Do(task)
	//s2.Every(2).Seconds().Do(task)
	//s2.Every(1).Minute().Do(task)
	//s2.Every(2).Minutes().Do(task)
	//s2.Every(1).Hour().Do(task)
	//s2.Every(2).Hours().Do(task)
	//s2.Every(1).Day().Do(task)
	//s2.Every(2).Days().Do(task)
	//s2.Every(1).Week().Do(task)
	//s2.Every(2).Weeks().Do(task)
	//s2.Every(1).Month(time.Now().Day()).Do(task)
	//s2.Every(2).Months(15).Do(task)
	//
	////// check for errors
	_, err := s2.Every(1).Day().At("01:09").Do(task)
	if err != nil {
		log.Fatalf("error creating job: %v", err)
	}
	//
	//// Do jobs with params
	//s2.Every(1).Second().Do(taskWithParams, 1, "hello")
	//
	//// Do Jobs with tags
	//// initialize tag
	//tag1 := []string{"tag1"}
	//tag2 := []string{"tag2"}
	//
	//s2.Every(1).Week().SetTag(tag1).Do(task)
	//s2.Every(1).Week().SetTag(tag2).Do(task)
	//
	//// Removing Job Based on Tag
	//s2.RemoveJobByTag("tag1")
	//
	//// Remove a Job after its last execution
	//j, _ := s2.Every(1).StartAt(time.Now().Add(30*time.Second)).Do(task)
	//j.LimitRunsTo(1)
	//j.RemoveAfterLastRun()

	//// Do jobs on specific weekday
	//s2.Every(1).Monday().Do(task)
	//s2.Every(1).Thursday().Do(task)

	// Do a job at a specific time - 'hour:min:sec' - seconds optional
	//s2.Every(1).Day().At("17:01:10").Do(task2)

	//_, err := s2.Cron("10 17 * * *").Do(task)
	//if err != nil {
	//	log.Println(err)
	//}
	//s2.Every(1).Monday().At("18:30").Do(task)
	//s2.Every(1).Tuesday().At("18:30:59").Do(task)
	//s2.Every(1).Wednesday().At("1:01").Do(task)

	// Begin job at a specific date/time.
	//t := time.Date(2021, time.June, 9, 20, 0, 0, 0, time.UTC)
	//s2.Every(1).Hour().StartAt(t).Do(task)

	// Delay start of job
	//s2.Every(1).Hour().StartAt(time.Now().Add(time.Duration(1 * time.Hour)).Do(task)

	// NextRun gets the next running time

	// Remove a specific job
	//s2.Remove(task)

	// Clear all scheduled jobs
	//s2.Clear()

	// stop our first scheduler (it still exists but doesn't run anymore)
	//s1.Stop()

	// executes the scheduler and blocks current thread
	s2.StartBlocking()
	//s2.StartAsync()
	// this line is never reached
}
