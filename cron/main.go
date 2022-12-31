package cron

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/DanielRenne/GoCore/core/cron"
	"github.com/DanielRenne/goCoreAppTemplate/scheduleEngine"
)

func Start() {
	defer func() {
		if r := recover(); r != nil {
			session_functions.Print("\n\nPanic Stack: " + string(debug.Stack()))
			session_functions.Log("main.go", "Panic Recovered at Start:  "+fmt.Sprintf("%+v", r))
			time.Sleep(time.Millisecond * 3000)
			Start()
			return
		}
	}()
	go cron.RegisterRecurring(cron.CronTopOf30Seconds, ClearLogs)
	go cron.RegisterRecurring(cron.CronTopOfHour, ClearDebugMemory)
	go cron.RegisterRecurring(cron.CronTopOfHour, DeleteImageHistory)
	go cron.RegisterRecurring(cron.CronTopOfSecond, FlushLogs)
	go cron.RegisterRecurring(cron.CronTopOfSecond, scheduleEngine.Trigger)
	go cron.RegisterRecurring(cron.CronTopOfSecond, BroadcastTime)
	go startup()
}
