package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uname    string `gorm:"embedded"`
	Password string
	Admin    bool
	Age      int
	FullName string
}
