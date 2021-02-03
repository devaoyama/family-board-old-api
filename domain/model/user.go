package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	LineUid    string `gorm:"unique;type:varchar(255)"`
	Name       string `gorm:"type:varchar(255)"`
	PictureUrl string `gorm:"type:text"`
}

func NewUser(lineUid, name string, pictureUrl string) *User {
	return &User{
		LineUid:    lineUid,
		Name:       name,
		PictureUrl: pictureUrl,
	}
}
