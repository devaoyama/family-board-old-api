package repository

import "family-board-api/domain/model"

type TodoRepository interface {
	FindById(id int) (*model.Todo, error)
	Store(todo *model.Todo) (*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
	Delete(todo *model.Todo) (*model.Todo, error)
}
