package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smoothlee/calendar/service"
)

func add(c *gin.Context) {
	token := c.PostForm("token")
	uid := service.GetUIDByToken(token)
	if uid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "token error",
		})
		return
	}
	year, _ := strconv.Atoi(c.PostForm("year"))
	month, _ := strconv.Atoi(c.PostForm("month"))
	day, _ := strconv.Atoi(c.PostForm("day"))
	time := c.PostForm("time")
	title := c.PostForm("title")
	if err := service.AddEvent(uid, year, month, day, time, title); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -2,
			"msg":  err.Error(),
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
	token := c.PostForm("token")
	uid := service.GetUIDByToken(token)
	if uid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "token error",
		})
		return
	}
	eventID := c.PostForm("event_id")
	if err := service.DelEvent(uid, eventID); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -2,
			"msg":  err.Error(),
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
	token := c.PostForm("token")
	uid := service.GetUIDByToken(token)
	if uid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "token error",
			"data": nil,
		})
		return
	}
	year, _ := strconv.Atoi(c.PostForm("year"))
	month, _ := strconv.Atoi(c.PostForm("month"))
	day, _ := strconv.Atoi(c.PostForm("day"))
	list, err := service.GetDay(uid, year, month, day)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": list,
	})
	return
}

func getMonth(c *gin.Context) {
	token := c.PostForm("token")
	fmt.Println(token)
	uid := service.GetUIDByToken(token)
	fmt.Println(uid)
	if uid == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "token error",
			"data": nil,
		})
		return
	}
	year, _ := strconv.Atoi(c.PostForm("year"))
	month, _ := strconv.Atoi(c.PostForm("month"))
	list, err := service.GetMonth(uid, year, month)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": list,
	})
	return
}
