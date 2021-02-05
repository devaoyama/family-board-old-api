package registry

import (
	"family-board-api/domain/repository"
	"family-board-api/infrastructure/persistence"

	"gorm.io/gorm"
)

type Repository struct {
	Ur repository.UserRepository
	Fr repository.FamilyRepository
}

func NewRepository(db *gorm.DB) *Repository {
	up := persistence.NewUserPersistence(db)
	fp := persistence.NewFamilyPersistence(db)

	return &Repository{
		Ur: up,
		Fr: fp,
	}
}
