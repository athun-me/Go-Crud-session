package main

import (
	"github.com/athunlal/config"
	"github.com/athunlal/controlls"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {

	config.Dbconnect()

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// err := r.Run(":3000")
	// if err != nil {
	//     panic("[Error] failed to start Gin server due to: " + err.Error())
	//     return
	// }

	r.POST("/signup", controlls.UserSignUP)
	r.POST("/login", controlls.UserLogin)

	r.POST("/logout", controlls.DeleteSession)
	r.GET("/login", controlls.Loginpage)
	r.GET("/admin", controlls.CheckAdmin, controlls.Adminpage)
	r.GET("/signup", controlls.SignUpPage)
	r.GET("/home", controlls.CheckSession, controlls.HomePage)
	r.POST("/delete", controlls.DeleteUser)

	r.Run(":4000")
}
