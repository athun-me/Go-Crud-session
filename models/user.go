package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uname    string `gorm:"embedded"`
	Password string
	Sub      string
	Admin    bool
}
