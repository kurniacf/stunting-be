package repository

import (
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type todoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) models.TodoRepositoryInterface {
	return &todoRepository{
		DB: db,
	}
}

func (r *todoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo

	err := r.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *todoRepository) FindById(id uint) (*models.Todo, error) {
	var todo models.Todo

	err := r.DB.First(&todo, id).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *todoRepository) Create(todo *models.Todo) (*models.Todo, error) {
	err := r.DB.Create(todo).Error
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *todoRepository) Update(todo *models.Todo) error {
	return r.DB.Clauses(clause.Returning{}).Save(todo).Error
}

func (r *todoRepository) Delete(todo *models.Todo) error {
	return r.DB.Delete(todo).Error
}
