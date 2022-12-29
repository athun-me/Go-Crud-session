package controlls

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("userId", 12090292)
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/login")
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "User Sign In successfully",
	// })
}

func CheckSession(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("userId") != nil {
		fmt.Println(session.Get("userId"))
		c.Redirect(http.StatusMovedPermanently, "/home")
		c.Next()
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
}

func DeleteSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Delete("userId")
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/login")
}
