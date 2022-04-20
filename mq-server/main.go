package main

import (
	"mq-server/conf"
	"mq-server/service"
)

/*
* @Author: hh
* @Date:   2022/4/18 19:42
 */
func main() {
	conf.Init()

	forerver:=make(chan bool)
	service.CreateTask()
	<-forerver
}
