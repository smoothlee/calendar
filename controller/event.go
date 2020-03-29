package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smoothlee/calendar/service"
)

func add(c *gin.Context) {
	uid := c.PostForm("uid")
	if !service.CheckUID(uid) {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "wrong uid please login",
		})
		return
	}
	event := c.PostForm("event")
	date := c.PostForm("date")
	if err := service.AddEvent(uid, date, event); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "OK",
	})
	return

}

func del(c *gin.Context) {
	uid := c.PostForm("uid")
	eventID := c.PostForm("event_id")
	if uid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "please login",
		})
	}
	if err := service.DelEvent(uid, eventID); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "OK",
	})
	return
}

func getDay(c *gin.Context) {
	uid := c.PostForm("uid")
	date := c.PostForm("date")
	if uid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "please login",
			"data": nil,
		})
	}
	eventList, err := service.GetEvents(uid, date)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "server error",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": eventList,
	})
	return
}

func getMonth(c *gin.Context) {

}
