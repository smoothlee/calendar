package controller

import "github.com/gin-gonic/gin"

func Init() {
	r := gin.Default()
	r.POST("/user/register", register)
	r.POST("/user/login", login)

	r.POST("/event/add", add)
	r.POST("/event/del", del)
	r.POST("/event/get_day", getDay)
	r.POST("/event/get_month", getMonth)
	r.Run(":8080")
}
