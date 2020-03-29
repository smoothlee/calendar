package main

import (
	"github.com/smoothlee/calendar/controller"
	"github.com/smoothlee/calendar/mysql"
	"github.com/smoothlee/calendar/service"
)

func Init() {
	if err := mysql.Init(); err != nil {
		panic(err)
	}
	service.Init()
	controller.Init()
}
