package config

import (
	"family-board-api/domain/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.User{}, &model.Family{}, &model.Todo{})

	return db, nil
}
