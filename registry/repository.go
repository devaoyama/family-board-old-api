package registry

import (
	"family-board-api/domain/repository"
	"family-board-api/infrastructure/persistence"

	"gorm.io/gorm"
)

type Repository struct {
	Ur repository.UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	up := persistence.NewUserPersistence(db)

	return &Repository{
		Ur: up,
	}
}
