package persistence

import (
	"family-board-api/domain/model"
	"family-board-api/domain/repository"

	"gorm.io/gorm"
)

type todoPersistence struct {
	Db *gorm.DB
}

func NewTodoPersistence(db *gorm.DB) repository.TodoRepository {
	return &todoPersistence{Db: db}
}

func (up *todoPersistence) FindById(id int) (*model.Todo, error) {
	var todo model.Todo
	result := up.Db.First(&todo, id)
	return &todo, result.Error
}

func (up *todoPersistence) Store(todo *model.Todo) (*model.Todo, error) {
	result := up.Db.Create(todo)
	return todo, result.Error
}

func (up *todoPersistence) Update(todo *model.Todo) (*model.Todo, error) {
	result := up.Db.Save(todo)
	return todo, result.Error
}

func (up *todoPersistence) Delete(todo *model.Todo) (*model.Todo, error) {
	result := up.Db.Delete(todo)
	return todo, result.Error
}
