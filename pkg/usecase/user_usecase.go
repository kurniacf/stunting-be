package usecase

import (
	"github.com/kurniacf/stunting-be/pkg/models"
)

type userUsecase struct {
	userRepo models.UserRepository
}

func NewUserUsecase(ur models.UserRepository) models.UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

// Here you can add methods to implement your use cases
