package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;unique" json:"username"`
	Password string `gorm:"size:255" json:"password"`
	Token    string `gorm:"size:255" json:"token"`
}
