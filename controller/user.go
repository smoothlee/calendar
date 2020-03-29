package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smoothlee/calendar/service"
)

func login(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	code, msg, token := service.Login(username, email, password)
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"token": token,
	})
	return

}

func register(c *gin.Context) {
	password := c.PostForm("password")
	if len(password) < 6 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "password too short",
		})
		return
	}
	username := c.PostForm("username")
	email := c.PostForm("email")
	//service.register
	code, msg, token := service.Register(username, email, password)
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"token": token,
	})
	return
}
