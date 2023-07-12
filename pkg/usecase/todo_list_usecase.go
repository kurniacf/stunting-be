package usecase

import (
	"errors"
	"time"

	"github.com/kurniacf/stunting-be/pkg/api"
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/gorm"
)

type todoListUseCase struct {
	todoListRepo models.TodoListRepositoryInterface
	childRepo    models.ChildRepositoryInterface
}

func NewTodoListUsecase(tr models.TodoListRepositoryInterface, cr models.ChildRepositoryInterface) models.TodoListUseCaseInterface {
	return &todoListUseCase{
		todoListRepo: tr,
		childRepo:    cr,
	}
}

func (u *todoListUseCase) FindAll() ([]models.TodoList, error) {
	todoLists, err := u.todoListRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return todoLists, nil
}

func (u *todoListUseCase) FindById(id uint) (*models.TodoList, error) {
	todo, err := u.todoListRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (u *todoListUseCase) FindByChildId(childId uint, userId uint, dateString string, done string) ([]models.TodoList, error) {
	var (
		date time.Time
		err  error
	)

	if dateString != "" {
		date, err = time.Parse("2006-01-02", dateString)
		if err != nil {
			return nil, errors.New("Date format invalid")
		}
	}

	child, err := u.childRepo.FindById(childId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Child data not found")
		}
		return nil, err
	}

	if child.UserID != userId {
		return nil, errors.New("You are not authorized to access this data")
	}

	todoLists, err := u.todoListRepo.FindByChildId(childId, date, done)
	if err != nil {
		return nil, err
	}

	return todoLists, nil
}

func (u *todoListUseCase) Create(request *api.CreateTodoListRequest) (*models.TodoList, error) {
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		return nil, errors.New("Birth date format invalid")
	}

	todoList := &models.TodoList{
		ChildID: request.ChildID,
		TodoID:  request.TodoID,
		Date:    date,
	}

	createdTodoList, err := u.todoListRepo.Create(todoList)
	if err != nil {
		return nil, err
	}

	return createdTodoList, nil
}

func (u *todoListUseCase) Update(id uint, request *api.CreateTodoListRequest) (*models.TodoList, error) {
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		return nil, errors.New("Birth date format invalid")
	}

	todoList, err := u.todoListRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Todo list data not found")
		}
		return nil, err
	}

	todoList.ChildID = request.ChildID
	todoList.TodoID = request.TodoID
	todoList.IsDone = request.IsDone
	todoList.Date = date

	err = u.todoListRepo.Update(*todoList)
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func (u *todoListUseCase) Delete(id uint) error {
	todoList, err := u.todoListRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("Todo list data not found")
		}
		return err
	}

	err = u.todoListRepo.Delete(todoList)
	if err != nil {
		return err
	}

	return nil
}
