package models

import (
	"time"

	"github.com/kurniacf/stunting-be/pkg/api"
	"gorm.io/gorm"
)

// User ...
type Child struct {
	gorm.Model
	Name         string    `gorm:"type:varchar(255);not null"`
	HealthStatus string    `gorm:"type:varchar(30)"`
	BirthDate    time.Time `gorm:"type:datetime;not null"`
	UserID       uint

	User User `gorm:"foreignKey:UserID"`
}

func (u *Child) TableName() string {
	return "childs"
}

type ChildUseCaseInterface interface {
	FindByUserId(uint) ([]Child, error)
	FindById(uint, uint) (*Child, error)
	Create(uint, *api.CreateChildRequest) (*Child, error)
	Update(uint, *api.CreateChildRequest) (*Child, error)
	Delete(uint, uint) error
}

type ChildRepositoryInterface interface {
	FindByUserId(uint) ([]Child, error)
	FindById(uint) (*Child, error)
	Create(*Child) (*Child, error)
	Update(*Child) error
	Delete(*Child) error
}
