package repository

import "family-board-api/domain/model"

type FamilyRepository interface {
	FindById(id int) (*model.Family, error)
	FindByInvitationCode(code string) (*model.Family, error)
	Store(family *model.Family) (*model.Family, error)
	Update(family *model.Family) (*model.Family, error)
	Delete(family *model.Family) (*model.Family, error)
	AppendTodo(family *model.Family, todo *model.Todo) (*model.Family,error)
}
