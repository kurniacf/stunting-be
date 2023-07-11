package usecase

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kurniacf/stunting-be/pkg/models"
	"time"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	user.Token = tokenString
	err = a.userUsecase.UpdateUser(user)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
