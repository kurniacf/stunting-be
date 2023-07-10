package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func InitDB() *gorm.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{})

	return db
}
