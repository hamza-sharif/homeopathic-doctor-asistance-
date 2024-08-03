package models

import "gorm.io/gorm"

type Medicine struct {
	gorm.Model
	Name        string `gorm:"size:255;not null" json:"name"`
	Description string `gorm:"size:255;" json:"description"`
}
