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

func (u *User) TableName() string {
	return "users"
}

type UserUsecase interface {
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id uint) error
	// add more methods as necessary
}

type UserRepository interface {
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll() ([]User, error)
	Save(user *User) error
	Update(user *User) error
	Delete(user *User) error
	// add more methods as necessary
}
