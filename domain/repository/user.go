package repository

import "family-board-api/domain/model"

type UserRepository interface {
	FindById(id int) (*model.User, error)
	FindByLineUid(lineUid string) (*model.User, error)
	Store(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(user *model.User) (*model.User, error)
	AppendFamily(user *model.User, family *model.Family) (*model.User, error)
}
