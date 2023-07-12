package usecase

import (
	"errors"

	"github.com/kurniacf/stunting-be/pkg/api"
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/gorm"
)

type todoUseCase struct {
	todoRepo models.TodoRepositoryInterface
}

func NewTodoUsecase(tr models.TodoRepositoryInterface) models.TodoUseCaseInterface {
	return &todoUseCase{
		todoRepo: tr,
	}
}

func (u *todoUseCase) FindAll() ([]models.Todo, error) {
	todos, err := u.todoRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (u *todoUseCase) FindById(id uint) (*models.Todo, error) {
	todo, err := u.todoRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (u *todoUseCase) Create(request *api.CreateTodoRequest) (*models.Todo, error) {
	todo := &models.Todo{
		Name: request.Name,
	}

	createdTodo, err := u.todoRepo.Create(todo)
	if err != nil {
		return nil, err
	}

	return createdTodo, nil
}

func (u *todoUseCase) Update(id uint, request *api.CreateTodoRequest) (*models.Todo, error) {
	todo, err := u.todoRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Todo data not found")
		}
		return nil, err
	}

	todo.Name = request.Name

	err = u.todoRepo.Update(todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (u *todoUseCase) Delete(id uint) error {
	todo, err := u.todoRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Todo data not found")
		}
		return err
	}

	err = u.todoRepo.Delete(todo)
	if err != nil {
		return err
	}

	return nil
}
