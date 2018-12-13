package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type cronjob struct {
	expr *cronexpr.Expression
	nextTime time.Time
}
func main(){
	var (
		task1, task2 *cronjob
		expr1, expr2 *cronexpr.Expression
		now time.Time
		scheduleTable map[string]*cronjob
	)

	scheduleTable = make(map[string]*cronjob)

	expr1 = cronexpr.MustParse("*/3 * * * * * *")
	task1 = &cronjob{
		expr:expr1,
		nextTime:expr1.Next(now),
	}
	scheduleTable["Task1"] = task1

	expr2 = cronexpr.MustParse("*/5 * * * * * *")
	task2 = &cronjob{
		expr:expr2,
		nextTime:expr2.Next(now),
	}
	scheduleTable["Task2"] = task2

	go func(){
		for{
			var(
				now time.Time
			)

			now = time.Now()
			for taskName, taskJob := range scheduleTable{
				if taskJob.nextTime.Before(now) || taskJob.nextTime.Equal(now){
					go func(name string){
						fmt.Println(name, "is execited at ", time.Now())
					}(taskName)
					taskJob.nextTime = taskJob.expr.Next(now)
				}
			}
		}

	}()

	select {
	case <- time.NewTimer(10 * time.Second).C:
	}


	time.Sleep(50*time.Second)
}