package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kurniacf/stunting-be/pkg/api"
	"github.com/kurniacf/stunting-be/pkg/middleware"
	"github.com/kurniacf/stunting-be/pkg/models"
)

type todoListHandler struct {
	todoListUseCase models.TodoListUseCaseInterface
}

func NewTodoListHandler(r *gin.RouterGroup, tu models.TodoListUseCaseInterface) {
	handler := &todoListHandler{
		todoListUseCase: tu,
	}

	r.GET("/", handler.FindAll)
	r.GET("/child/:id", middleware.JwtAuthMiddleware(), handler.FindByChildId)
	r.GET("/:id", handler.FindById)
	r.POST("/", handler.Create)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
}

func (h *todoListHandler) FindAll(c *gin.Context) {
	todoLists, err := h.todoListUseCase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	todoListResponses := []api.TodoListResponse{}
	for _, todoList := range todoLists {
		todoListResponses = append(todoListResponses, api.TodoListResponse{
			ID:      todoList.ID,
			IsDone:  todoList.IsDone,
			Date:    todoList.Date,
			ChildID: todoList.ChildID,
			TodoID:  todoList.TodoID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todoListResponses,
	})
}

func (h *todoListHandler) FindById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID invalid",
		})
		return
	}

	todoList, err := h.todoListUseCase.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": api.TodoListResponse{
			ID:      todoList.ID,
			IsDone:  todoList.IsDone,
			Date:    todoList.Date,
			ChildID: todoList.ChildID,
			TodoID:  todoList.TodoID,
			ToDo: &api.TodoResponse{
				ID:   todoList.Todo.ID,
				Name: todoList.Todo.Name,
			},
		},
	})
}

func (h *todoListHandler) FindByChildId(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token invalid",
		})
		return
	}

	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID invalid",
		})
		return
	}

	todoLists, err := h.todoListUseCase.FindByChildId(uint(id), uint(userId.(float64)), c.Query("date"), c.Query("done"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	todoListResponses := []api.TodoListResponse{}
	for _, todoList := range todoLists {
		todoListResponses = append(todoListResponses, api.TodoListResponse{
			ID:      todoList.ID,
			IsDone:  todoList.IsDone,
			Date:    todoList.Date,
			ChildID: todoList.ChildID,
			TodoID:  todoList.TodoID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todoListResponses,
	})
}

func (h *todoListHandler) Create(c *gin.Context) {
	var req api.CreateTodoListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todoList, err := h.todoListUseCase.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": api.TodoListResponse{
			ID:      todoList.ID,
			IsDone:  todoList.IsDone,
			Date:    todoList.Date,
			ChildID: todoList.ChildID,
			TodoID:  todoList.TodoID,
			Child:   nil,
			ToDo:    nil,
		},
	})
}

// TODO: foreign key todo_id and child_id can't be updated
func (h *todoListHandler) Update(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID invalid",
		})
		return
	}

	var req api.CreateTodoListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todoList, err := h.todoListUseCase.Update(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": api.TodoListResponse{
			ID:      todoList.ID,
			IsDone:  todoList.IsDone,
			Date:    todoList.Date,
			ChildID: todoList.ChildID,
			TodoID:  todoList.TodoID,
		},
	})
}

func (h *todoListHandler) Delete(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID invalid",
		})
		return
	}

	err = h.todoListUseCase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
