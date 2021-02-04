package persistence

import (
	"family-board-api/domain/model"
	"family-board-api/domain/repository"

	"gorm.io/gorm"
)

type familyPersistence struct {
	Db *gorm.DB
}

func NewFamilyPersistence(db *gorm.DB) repository.FamilyRepository {
	return &familyPersistence{Db: db}
}

func (fp *familyPersistence) FindById(id int) (*model.Family, error) {
	var family model.Family
	result := fp.Db.First(&family, id)
	return &family, result.Error
}

func (fp *familyPersistence) Store(family *model.Family) (*model.Family, error) {
	result := fp.Db.Create(family)
	return family, result.Error
}

func (fp *familyPersistence) Update(family *model.Family) (*model.Family, error) {
	result := fp.Db.Save(family)
	return family, result.Error
}

func (fp *familyPersistence) Delete(family *model.Family) (*model.Family, error) {
	result := fp.Db.Delete(family)
	return family, result.Error
}