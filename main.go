package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

var taipeiZone *time.Location

func init() {
	taipeiZone, _ = time.LoadLocation("Asia/Taipei")
}

func main() {
	start := time.Now().In(taipeiZone).Format("2006/01/02 15:04:05")
	fmt.Println("Start at", start)

	// 建立排程
	cronjob := cron.New(cron.WithSeconds())

	// 每兩秒執行一次動作
	cronjob.AddFunc("*/2 * * * * *", func() {
		printCronJobProgessBySec(2)
	})

	// 每五秒執行一次動作
	cronjob.AddFunc("*/5 * * * * *", func() {
		printCronJobProgessBySec(5)
	})

	// 排程開始
	cronjob.Start()

	time.Sleep(30 * time.Second)

	end := time.Now().In(taipeiZone).Format("2006/01/02 15:04:05")
	fmt.Println("End at", end)
}

// printCronJobProgessBySec 透過秒數打印排程進度
func printCronJobProgessBySec(sec int) {
	now := time.Now().In(taipeiZone).Format("2006/01/02 15:04:05")
	fmt.Println("Cron job: do it every", sec, "seconds. Time:", now)
}
