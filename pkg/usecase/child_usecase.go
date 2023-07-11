package usecase

import (
	"errors"
	"time"

	"github.com/kurniacf/stunting-be/pkg/api"
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/gorm"
)

type childUsecase struct {
	childRepo models.ChildRepositoryInterface
}

func NewChildUsecase(cr models.ChildRepositoryInterface) models.ChildUseCaseInterface {
	return &childUsecase{
		childRepo: cr,
	}
}

func (u *childUsecase) FindByUserId(userId uint) ([]models.Child, error) {
	childs, err := u.childRepo.FindByUserId(userId)
	if err != nil {
		return nil, err
	}

	return childs, nil
}

func (u *childUsecase) FindById(userId uint, id uint) (*models.Child, error) {
	child, err := u.childRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	if child.UserID != userId {
		return nil, errors.New("You are not authorized to this data")
	}

	return child, nil
}

func (u *childUsecase) Create(userId uint, request *api.CreateChildRequest) (*models.Child, error) {
	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	if err != nil {
		return nil, errors.New("Birth date format invalid")
	}

	child := &models.Child{
		Name:         request.Name,
		HealthStatus: request.HealthStatus,
		BirthDate:    birthDate,
		UserID:       userId,
	}

	createdChild, err := u.childRepo.Create(child)
	if err != nil {
		return nil, err
	}

	return createdChild, nil
}

func (u *childUsecase) Update(id uint, request *api.CreateChildRequest) (*models.Child, error) {
	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	if err != nil {
		return nil, errors.New("Birth date format invalid")
	}

	child, err := u.childRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Child data not found")
		}
		return nil, err
	}

	child.Name = request.Name
	child.HealthStatus = request.HealthStatus
	child.BirthDate = birthDate

	err = u.childRepo.Update(child)
	if err != nil {
		return nil, err
	}

	return child, nil
}

func (u *childUsecase) Delete(userId uint, id uint) error {
	child, err := u.childRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Child data not found")
		}
		return err
	}

	if child.UserID != userId {
		return errors.New("You are not authorized to this data")
	}

	err = u.childRepo.Delete(child)
	if err != nil {
		return err
	}

	return nil
}
