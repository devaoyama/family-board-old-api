package registry

import (
	"family-board-api/domain/repository"
	"family-board-api/infrastructure/persistence"

	"gorm.io/gorm"
)

type repo struct {
	User repository.UserRepository
}

func NewRepository(db *gorm.DB) *repo {
	up := persistence.NewUserPersistence(db)

	return &repo{
		User: up,
	}
}
