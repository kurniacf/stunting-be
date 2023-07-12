package seeds

import (
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/gorm"
	"time"
)

func SeedTodoLists(db *gorm.DB) error {
	// Fetch all children and todos
	var children []models.Child
	db.Find(&children)

	var todos []models.Todo
	db.Find(&todos)

	todoLists := []models.TodoList{}

	for _, child := range children {
		for _, todo := range todos {
			todoList := models.TodoList{
				ChildID: child.ID,
				TodoID:  todo.ID,
				IsDone:  false,
				Date:    time.Now(),
			}
			todoLists = append(todoLists, todoList)
		}
	}

	for _, todoList := range todoLists {
		if err := db.Create(&todoList).Error; err != nil {
			return err
		}
	}

	return nil
}
