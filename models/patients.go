package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID                  string `gorm:"primaryKey" json:"id"`
	Name                string `gorm:"size:100" json:"name"`
	Address             string `gorm:"size:255" json:"address"`
	Age                 int    `gorm:"default:0" json:"age"`
	Cnic                string `gorm:"size:50" json:"cnic"`
	FatherOrHusbandName string `gorm:"size:50" json:"fatherOrHusbandName"`
	Gender              string `gorm:"size:4" json:"gender"`
	MobileNo            string `gorm:"size:20" json:"mobileNo"`
	Weight              int    `gorm:"default:0" json:"weight"`
	Price               int    `gorm:"default:0" json:"price"`
	Medicine            string `gorm:"size:100" json:"medicine"`
}

type Price struct {
	gorm.Model
	Fee int `gorm:"default:0" json:"fee"`
}
