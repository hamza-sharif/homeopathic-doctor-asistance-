package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name                string  `gorm:"size:255" json:"name"`
	Address             string  `gorm:"size:255" json:"address"`
	Age                 float32 `gorm:"default:0" json:"age"`
	Cnic                string  `gorm:"size:255" json:"cnic"`
	FatherOrHusbandName string  `gorm:"size:255" json:"fatherOrHusbandName"`
	Gender              string  `gorm:"size:4" json:"gender"`
	MobileNo            string  `gorm:"size:255" json:"mobileNo"`
	Weight              int     `gorm:"default:0" json:"weight"`
}
