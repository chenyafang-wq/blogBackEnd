package utils

import (
	"github.com/robfig/cron/v3"
	"log"
)

func TimingSpider(cmd func()) {

	log.Println("cron TimingSpider start:")

	// v3 用法 干
	c := cron.New(cron.WithSeconds())

	// 每天定时执行的条件
	//spec := viper.GetString(`cron.timing_spider`)
	c.AddFunc("@every 1s", func() {
		cmd()
	})

	go c.Start()

	// 关闭计划任务, 但是不能关闭已经在执行中的任务.
	// defer c.Stop()

	// 阻塞
	//select {}
}

func RecentUpdate(cmd func()) {
	c := cron.New()
	// 每天定时执行的条件
	spec := "* * * * *"

	c.AddFunc(spec, func() {
		cmd()
	})

	go c.Start()
}
