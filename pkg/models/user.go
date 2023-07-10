package models

import (
	"gorm.io/gorm"
)

// User ...
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"unique;type:varchar(255);not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

type UserUsecase interface {
	// define your usecase methods here
}

type UserRepository interface {
	// define your repository methods here
}
