package main

import (
	"github.com/smoothlee/calendar/controller"
	"github.com/smoothlee/calendar/service"
)

func Init() {
	mysql.Init()
	service.Init()
	controller.Init()
}
