package configs

import (
	"fmt"
	"os"

	"github.com/kurniacf/stunting-be/pkg/seeds"

	"github.com/joho/godotenv"
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(seed bool, prod bool) *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}

	var dsn string
	if prod {
		dsn = os.Getenv("DSN")
	} else {
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(&models.User{}, &models.Child{}, &models.Todo{}, &models.TodoList{})
	if err != nil {
		panic("Failed to perform database migration")
	}

	if seed {
		err = seeds.SeedUsers(db)
		if err != nil {
			panic("Failed to seed users")
		}
		err = seeds.SeedChildren(db)
		if err != nil {
			panic("Failed to seed children")
		}
		err = seeds.SeedTodos(db)
		if err != nil {
			panic("Failed to seed todos")
		}
		err = seeds.SeedTodoLists(db)
		if err != nil {
			panic("Failed to seed todo lists")
		}
	}

	return db
}
