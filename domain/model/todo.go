package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	FamilyId    int
	Title       string    `gorm:"type:varchar(255)"`
	Description string    `gorm:"type:varchar(255)"`
	Status      bool      `gorm:"type:bool;default:false"`
	Date        time.Time `gorm:"type:date"`
}

func NewTodo(title, description string, status bool, date time.Time) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Status:      status,
		Date:        date,
	}
}
