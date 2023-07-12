package models

import (
	"github.com/kurniacf/stunting-be/pkg/api"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string `gorm:"size:100;not null;unique"`
}

func (u *Todo) TableName() string {
	return "todos"
}

type TodoUseCaseInterface interface {
	FindAll() ([]Todo, error)
	FindById(uint) (*Todo, error)
	Create(*api.CreateTodoRequest) (*Todo, error)
	Update(uint, *api.CreateTodoRequest) (*Todo, error)
	Delete(uint) error
}

type TodoRepositoryInterface interface {
	FindAll() ([]Todo, error)
	FindById(uint) (*Todo, error)
	Create(*Todo) (*Todo, error)
	Update(*Todo) error
	Delete(*Todo) error
}
