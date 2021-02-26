package repository

import (
	"family-board-api/domain/model"
	"time"
)

type TodoRepository interface {
	FindById(id int) (*model.Todo, error)
	FindByFamilyId(familyId int) ([]*model.Todo, error)
	FindByFamilyIdAndDate(familyId int, date time.Time) ([]*model.Todo, error)
	Store(todo *model.Todo) (*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
	Delete(todo *model.Todo) (*model.Todo, error)
}
