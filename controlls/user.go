package controlls

import (
	
	"html/template"
	"log"
	"net/http"

	"github.com/athunlal/config"
	"github.com/athunlal/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserSignUP(c *gin.Context) {
	c.Request.ParseForm()
	type Data struct {
		Uname    string
		Password string
		Sub    string
	}

	var data Data

	data.Uname = c.Request.PostForm["uname"][0]
	data.Sub = c.Request.PostForm["sub"][0]
	data.Password = c.Request.PostForm["password"][0]

	user := models.User{Uname: data.Uname, Password: data.Password, Sub: data.Uname}

	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "User not inserted",
		})
		return
	}

	// c.JSON(200, gin.H{
	// 	"message": "Sign up succesfull",
	// })

	c.Redirect(http.StatusMovedPermanently, "/login")
}

func Loginpage(c *gin.Context) {
	//html page integration
	tmpl := template.Must(template.ParseFiles("template/login.html"))
	tmpl.Execute(c.Writer, nil)

}

func SignUpPage(c *gin.Context) {
	//html page integration
	tmpl := template.Must(template.ParseFiles("template/signup.html"))
	tmpl.Execute(c.Writer, nil)

}

func HomePage(c *gin.Context) {
	//html page integration
	tmpl := template.Must(template.ParseFiles("template/home.html"))
	tmpl.Execute(c.Writer, nil)

}

func DeletUser(c *gin.Context){

}

func Adminpage(c *gin.Context) {
	var users []models.User

	result := config.DB.Find(&users)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Users not found",
		})
		return
	}

	tmpl := template.Must(template.ParseFiles("template/admin.html"))
	tmpl.Execute(c.Writer, users)
}

func UserLogin(c *gin.Context) {

	c.Request.ParseForm()

	type Data struct {
		Uname    string
		Password string
		Admin    bool
	}

	var data Data

	data.Uname = c.Request.PostForm["uname"][0]
	data.Password = c.Request.PostForm["password"][0]

	var user models.User

	result := config.DB.First(&user, "uname = ?", data.Uname)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "User not found",
		})
		return
	}

	if user.Password == data.Password {
		session := sessions.Default(c)
		session.Set("userId", user.ID)
		session.Save()

		// c.JSON(200, gin.H{
		// 	"data": user,
		// })

		log.Fatal(user.Admin)
		if user.Admin == true {
			c.Redirect(http.StatusMovedPermanently, "/admin")
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/home")
		return
	}

	c.JSON(500, gin.H{
		"message": "Not found",
	})
}
