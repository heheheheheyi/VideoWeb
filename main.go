package main

import (
	"VideoWeb/config"
	"VideoWeb/model"
	"VideoWeb/router"
	"VideoWeb/task"
)

func main() {
	config.Init()
	model.InitMysql()
	model.InitRedis()
	task.CronJob()
	router.InitRouter()
}
