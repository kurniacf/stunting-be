package repository

import (
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type childRepository struct {
	DB *gorm.DB
}

func NewChildRepository(db *gorm.DB) models.ChildRepositoryInterface {
	return &childRepository{
		DB: db,
	}
}

func (r *childRepository) FindByUserId(userId uint) ([]models.Child, error) {
	var childs []models.Child
	err := r.DB.Where("user_id", userId).Find(&childs).Error
	if err != nil {
		return nil, err
	}
	return childs, nil
}

func (r *childRepository) FindById(id uint) (*models.Child, error) {
	var child models.Child
	err := r.DB.First(&child, id).Error
	if err != nil {
		return nil, err
	}
	return &child, nil
}

func (r *childRepository) Create(child *models.Child) (*models.Child, error) {
	err := r.DB.Create(child).Error
	if err != nil {
		return nil, err
	}

	return child, nil
}

func (r *childRepository) Update(child *models.Child) error {
	return r.DB.Clauses(clause.Returning{}).Save(child).Error
}

func (r *childRepository) Delete(child *models.Child) error {
	return r.DB.Delete(child).Error
}
