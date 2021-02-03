package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	LineUid string `gorm:"unique;type:varchar(255)"`
	Name    string `gorm:"type:varchar(255)"`
}

func NewUser(lineUid, name string) *User {
	return &User{
		LineUid: lineUid,
		Name:    name,
	}
}
