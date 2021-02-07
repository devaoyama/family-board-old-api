package usecase

import (
	"errors"
	"family-board-api/domain/model"
	"family-board-api/domain/repository"
	"family-board-api/usecase/input"
	"family-board-api/usecase/output"
	"time"
)

type TodoUsecase struct {
	ur repository.UserRepository
	fr repository.FamilyRepository
	tr repository.TodoRepository
}

func NewTodoUsecase(ur repository.UserRepository, fr repository.FamilyRepository, tr repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		ur: ur,
		fr: fr,
		tr: tr,
	}
}

func (tu *TodoUsecase) CreateTodo(i *input.CreateTodo) (*output.CreateTodo, error) {
	o := &output.CreateTodo{}
	user, err := tu.ur.FindById(i.UserId)
	if err != nil {
		return o, nil
	}
	todo := model.NewTodo(i.Title, i.Description, false, time.Now())
	family, err := tu.fr.FindById(*user.FamilyId)
	if err != nil {
		return o, err
	}
	family, err = tu.fr.AppendTodo(family, todo)
	o.Todo = todo

	return o, err
}

func (tu *TodoUsecase) ChangeTodoStatus(i *input.ChangeTodoStatus) (*output.ChangeTodoStatus, error) {
	o := &output.ChangeTodoStatus{}
	user, err := tu.ur.FindById(i.UserId)
	if err != nil {
		return o, nil
	}
	todo, err := tu.tr.FindById(i.TodoId)
	if err != nil {
		return o, nil
	}
	if *user.FamilyId != todo.FamilyId {
		return o, errors.New("this action is unauthorized")
	}
	todo.Status = !todo.Status
	todo, err = tu.tr.Update(todo)
	o.Todo = todo
	return o, nil
}

func (tu *TodoUsecase) DeleteTodo(i *input.DeleteTodo) (*output.DeleteTodo, error) {
	o := &output.DeleteTodo{}
	user, err := tu.ur.FindById(i.UserId)
	if err != nil {
		return o, nil
	}
	todo, err := tu.tr.FindById(i.TodoId)
	if err != nil {
		return o, nil
	}
	if *user.FamilyId != todo.FamilyId {
		return o, errors.New("this action is unauthorized")
	}
	todo, err = tu.tr.Delete(todo)
	o.Todo = todo
	return o, err
}
