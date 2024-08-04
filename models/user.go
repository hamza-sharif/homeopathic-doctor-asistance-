package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primarykey"`
	Username string `gorm:"size:255;unique" json:"username"`
	Password string `gorm:"size:255" json:"password"`
	Token    string `gorm:"size:255" json:"token"`
}

func (u *User) ConvertTOMap() map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(u)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}
