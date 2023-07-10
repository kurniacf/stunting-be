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

func (r *mysqlUserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mysqlUserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mysqlUserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *mysqlUserRepository) Save(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *mysqlUserRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *mysqlUserRepository) Delete(user *models.User) error {
	return r.DB.Delete(user).Error
}
