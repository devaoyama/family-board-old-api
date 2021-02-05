package persistence

import (
	"family-board-api/domain/model"
	"family-board-api/domain/repository"

	"gorm.io/gorm"
)

type userPersistence struct {
	Db *gorm.DB
}

func NewUserPersistence(db *gorm.DB) repository.UserRepository {
	return &userPersistence{Db: db}
}

func (up *userPersistence) FindById(id int) (*model.User, error) {
	var user model.User
	result := up.Db.First(&user, id)
	return &user, result.Error
}

func (up *userPersistence) FindByLineUid(lineUid string) (*model.User, error) {
	var user model.User
	result := up.Db.Where("line_uid", lineUid).First(&user)
	return &user, result.Error
}

func (up *userPersistence) Store(user *model.User) (*model.User, error) {
	result := up.Db.Create(user)
	return user, result.Error
}

func (up *userPersistence) Update(user *model.User) (*model.User, error) {
	result := up.Db.Save(user)
	return user, result.Error
}

func (up *userPersistence) Delete(user *model.User) (*model.User, error) {
	result := up.Db.Delete(user)
	return user, result.Error
}

func (up *userPersistence) AppendFamily(user *model.User, family *model.Family) (*model.User, error) {
	err := up.Db.Model(user).Association("Families").Append(family)
	return user, err
}
