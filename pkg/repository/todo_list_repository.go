package repository

import (
	"time"

	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type todoListRepository struct {
	DB *gorm.DB
}

func NewTodoListRepository(db *gorm.DB) models.TodoListRepositoryInterface {
	return &todoListRepository{
		DB: db,
	}
}

func (r *todoListRepository) FindAll() ([]models.TodoList, error) {
	var todoLists []models.TodoList

	err := r.DB.Find(&todoLists).Error
	if err != nil {
		return nil, err
	}

	return todoLists, nil
}

func (r *todoListRepository) FindById(id uint) (*models.TodoList, error) {
	var todoList models.TodoList

	err := r.DB.Preload("Todo").Where("id", id).First(&todoList).Error
	if err != nil {
		return nil, err
	}

	return &todoList, nil
}

func (r *todoListRepository) FindByChildId(childId uint, date time.Time, done string) ([]models.TodoList, error) {
	var todoList []models.TodoList

	tx := r.DB.Where("child_id = ?", childId)
	if !date.IsZero() {
		tx = tx.Where("date = ?", date)
	}

	if done == "true" {
		tx = tx.Where("is_done = ?", true)
	} else if done == "false" {
		tx = tx.Where("is_done = ?", false)
	}

	err := tx.Find(&todoList).Error
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func (r *todoListRepository) Create(todoList *models.TodoList) (*models.TodoList, error) {
	err := r.DB.Create(todoList).Error
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func (r *todoListRepository) Update(todoList models.TodoList) error {
	return r.DB.Clauses(clause.Returning{}).Save(&todoList).Error
}

func (r *todoListRepository) Delete(todoList *models.TodoList) error {
	return r.DB.Delete(todoList).Error
}
