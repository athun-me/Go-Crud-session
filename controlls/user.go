package controlls

import (
	"html/template"
	"net/http"
	"strconv"

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
		Age      string
		FullName string
	}

	var data Data

	data.Uname = c.Request.PostForm["uname"][0]
	data.Password = c.Request.PostForm["password"][0]
	data.FullName = c.Request.PostForm["fullname"][0]
	data.Age = c.Request.PostForm["age"][0]
	Int, _ := strconv.Atoi(data.Age)
	if data.Uname == "" || data.Password == "" || data.FullName == "" || data.Age == "" {
		c.Redirect(http.StatusMovedPermanently, "/signup")
		return
	}
	user := models.User{Uname: data.Uname, Password: data.Password, FullName: data.FullName, Age: Int}

	result := config.DB.Create(&user)

	if result.Error != nil {
		c.Redirect(http.StatusMovedPermanently, "/signup")
		return
	}

	// c.JSON(200, gin.H{
	// 	"message": "Sign up succesfull",
	// })

	c.Redirect(http.StatusMovedPermanently, "/login")
}

func Loginpage(c *gin.Context) {
	//html page integration
	session := sessions.Default(c)
	errMessage := session.Get("err")
	tmpl := template.Must(template.ParseFiles("template/login.html"))
	tmpl.Execute(c.Writer, errMessage)

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

func DeletUser(c *gin.Context) {

}

func Adminpage(c *gin.Context) {
	var users []models.User

	result := config.DB.Find(&users)

	if result.Error != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	tmpl := template.Must(template.ParseFiles("template/admin.html"))
	tmpl.Execute(c.Writer, users)
}

func UserLogin(c *gin.Context) {
	session := sessions.Default(c)
	c.Request.ParseForm()
	session.Clear()
	session.Save()
	type Data struct {
		Uname    string
		Password string
		Admin    bool
	}

	var data Data

	data.Uname = c.Request.PostForm["uname"][0]
	data.Password = c.Request.PostForm["password"][0]
	if data.Password == "" || data.Uname == "" {
		errorMessage := "username or password not found !"
		session.Set("err", errorMessage)
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	var user models.User

	result := config.DB.First(&user, "uname = ?", data.Uname)

	if result.Error != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}

	if user.Password == data.Password {
		session.Set("userId", user.ID)
		session.Save()
		if user.Admin == true {
			c.Redirect(http.StatusMovedPermanently, "/admin")
			return
		} else {
			c.Redirect(http.StatusMovedPermanently, "/home")
			return
		}

	}

	c.JSON(500, gin.H{
		"message": "Not found",
	})
}

func DeleteUser(c *gin.Context) {
	c.Request.ParseForm()
	uid := c.Request.PostForm["id"][0]
	config.DB.Delete(&models.User{}, uid)
	c.Redirect(http.StatusMovedPermanently, "/admin")
	return
}
