package repository

import (
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) models.UserRepository {
	return &mysqlUserRepository{
		DB: db,
	}
}

// Here you can add methods to interact with the DB
