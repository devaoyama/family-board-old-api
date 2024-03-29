package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FamilyId   *int
	LineUid    string `gorm:"unique;type:varchar(255)" json:"-"`
	Name       string `gorm:"type:varchar(255)"`
	PictureUrl string `gorm:"type:text"`
	Status     string `gorm:"type:varchar(255)"`
}

func NewUser(lineUid, name, pictureUrl string) *User {
	return &User{
		LineUid:    lineUid,
		Name:       name,
		PictureUrl: pictureUrl,
	}
}
