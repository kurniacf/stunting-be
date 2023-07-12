package seeds

import (
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/gorm"
	"time"
)

func SeedChildren(db *gorm.DB) error {
	// fetch all users
	var users []models.User
	db.Find(&users)

	children := []models.Child{
		{Name: "Peb JR", HealthStatus: "Sehat", BirthDate: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC), UserID: users[0].ID},
		{Name: "Damian Widodo", HealthStatus: "Sehat", BirthDate: time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC), UserID: users[1].ID},
		{Name: "Nur D. Pamungkas", HealthStatus: "Sehat", BirthDate: time.Date(2019, time.March, 3, 0, 0, 0, 0, time.UTC), UserID: users[2].ID},
		{Name: "Candra Prabowo", HealthStatus: "Kurang Sehat", BirthDate: time.Date(2018, time.April, 4, 0, 0, 0, 0, time.UTC), UserID: users[3].ID},
		{Name: "Puan Rani Septha Utami", HealthStatus: "Sehat", BirthDate: time.Date(2017, time.May, 5, 0, 0, 0, 0, time.UTC), UserID: users[4].ID},
	}

	for _, child := range children {
		if err := db.Create(&child).Error; err != nil {
			return err
		}
	}

	return nil
}
