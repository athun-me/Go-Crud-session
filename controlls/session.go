package controlls

import (
	"fmt"
	"log"
	"net/http"

	"github.com/athunlal/config"
	"github.com/athunlal/models"
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
		c.Next()
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
}

func DeleteSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	if session.Save() != nil {
		log.Fatal("Not deleted session")
	}

	c.Redirect(http.StatusMovedPermanently, "/login")
}

func CheckAdmin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("userId") != nil {
		fmt.Println(session.Get("userId"))
		var user models.User
		result := config.DB.First(&user, "ID = ?", session.Get("userId"))
		if result.Error != nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}
		if user.Admin == true {
			c.Next()
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
}
