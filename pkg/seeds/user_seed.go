package seeds

import (
	"github.com/kurniacf/stunting-be/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	users := []models.User{
		{Name: "Kurnia Cahya", Email: "kurniacf@gmail.com", Password: "password1"},
		{Name: "Marcel Mamahit", Email: "marcelmamahit@gmail.com", Password: "password2"},
		{Name: "Nur Muhammad", Email: "nurmuh@gmai.com", Password: "password3"},
		{Name: "Putri Salma", Email: "putrisalma@yahoo.com", Password: "password4"},
		{Name: "Lala Lulu", Email: "lulalalu@gmail.com", Password: "password5"},
	}

	for _, user := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)

		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
