package usecase

import (
	"github.com/kurniacf/stunting-be/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo models.UserRepository
}

func NewUserUsecase(ur models.UserRepository) models.UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (u *userUsecase) GetByID(id uint) (*models.User, error) {
	return u.userRepo.FindByID(id)
}

func (u *userUsecase) GetByEmail(email string) (*models.User, error) {
	return u.userRepo.FindByEmail(email)
}

func (u *userUsecase) GetAll() ([]models.User, error) {
	return u.userRepo.FindAll()
}

func (u *userUsecase) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return u.userRepo.Save(user)
}

func (u *userUsecase) UpdateUser(user *models.User) error {
	return u.userRepo.Update(user)
}

func (u *userUsecase) DeleteUser(id uint) error {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	return u.userRepo.Delete(user)
}

func (u *userUsecase) CheckPassword(email string, password string) error {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
