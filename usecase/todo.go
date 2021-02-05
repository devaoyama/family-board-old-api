package usecase

import (
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
	todo := model.NewTodo(i.Title, i.Description, false, time.Now())
	family, err := tu.fr.FindById(i.FamilyId)
	if err != nil {
		return o, err
	}
	// todo UserがFamilyに入っているかを確認
	family, err = tu.fr.AppendTodo(family, todo)
	o.Todo = todo

	return o, err
}

func (tu *TodoUsecase) ChangeTodoStatus(i *input.ChangeTodoStatus) (*output.ChangeStatus, error) {
	o := &output.ChangeStatus{}
	todo, err := tu.tr.FindById(i.TodoId)
	if err != nil {
		return o, nil
	}
	// todo UserがFamilyに入っているかを確認
	todo.Status = !todo.Status
	todo, err = tu.tr.Update(todo)
	o.Todo = todo
	return o, nil
}
