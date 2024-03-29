package model

import "gorm.io/gorm"

type Family struct {
	gorm.Model
	Todos          []*Todo
	Users          []*User
	Name           string `gorm:"type:varchar(255)"`
	InvitationCode string `gorm:"type:varchar(255)"`
}

func NewFamily(name, invitationCode string) *Family {
	return &Family{
		Name:           name,
		InvitationCode: invitationCode,
	}
}
