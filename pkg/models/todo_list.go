package models

import (
	"time"

	"github.com/kurniacf/stunting-be/pkg/api"
	"gorm.io/gorm"
)

type TodoList struct {
	gorm.Model
	ChildID uint      `gorm:"uniqueIndex:idx_todolist"`
	TodoID  uint      `gorm:"uniqueIndex:idx_todolist"`
	IsDone  bool      `gorm:"default:false"`
	Date    time.Time `gorm:"type:datetime;not null"`

	Child Child `gorm:"foreignKey:ChildID"`
	Todo  Todo  `gorm:"foreignKey:TodoID"`
}

func (u *TodoList) TableName() string {
	return "todo_lists"
}

type TodoListUseCaseInterface interface {
	FindAll() ([]TodoList, error)
	FindById(uint) (*TodoList, error)
	FindByChildId(uint, uint, string, string) ([]TodoList, error)
	Create(*api.CreateTodoListRequest) (*TodoList, error)
	Update(uint, *api.CreateTodoListRequest) (*TodoList, error)
	Delete(uint) error
}

type TodoListRepositoryInterface interface {
	FindAll() ([]TodoList, error)
	FindById(uint) (*TodoList, error)
	FindByChildId(uint, time.Time, string) ([]TodoList, error)
	Create(*TodoList) (*TodoList, error)
	Update(TodoList) error
	Delete(*TodoList) error
}
