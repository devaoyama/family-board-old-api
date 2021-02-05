package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Families   []*Family `gorm:"many2many:user_family;"`
	LineUid    string    `gorm:"unique;type:varchar(255)"`
	Name       string    `gorm:"type:varchar(255)"`
	PictureUrl string    `gorm:"type:text"`
}

func NewUser(lineUid, name, pictureUrl string) *User {
	return &User{
		LineUid:    lineUid,
		Name:       name,
		PictureUrl: pictureUrl,
	}
}
