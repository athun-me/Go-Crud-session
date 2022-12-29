package config

import (
	"github.com/athunlal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Dbconnect() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/app?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.User{})

}
