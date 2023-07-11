package usecase

import (
	"github.com/kurniacf/stunting-be/pkg/helper"
	"github.com/kurniacf/stunting-be/pkg/models"
)

type AuthUsecase interface {
	Login(email string, password string) (string, error)
}

type authUsecase struct {
	userUsecase models.UserUsecase
}

func NewAuthUsecase(userUsecase models.UserUsecase) AuthUsecase {
	return &authUsecase{
		userUsecase: userUsecase,
	}
}

func (a *authUsecase) Login(email string, password string) (string, error) {
	user, err := a.userUsecase.GetByEmail(email)
	if err != nil {
		return "", err
	}

	err = a.userUsecase.CheckPassword(user.Email, password)
	if err != nil {
		return "", err
	}

	token, err := helper.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	user.Token = token
	err = a.userUsecase.UpdateUser(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
