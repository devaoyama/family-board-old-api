package model

import "gorm.io/gorm"

type Family struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
}

func NewFamily(name string) *Family {
	return &Family{
		Name: name,
	}
}
