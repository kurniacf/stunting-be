package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kurniacf/stunting-be/pkg/api"
	"github.com/kurniacf/stunting-be/pkg/models"
)

type todoHandler struct {
	todoUseCase models.TodoUseCaseInterface
}

func NewTodoHandler(r *gin.RouterGroup, tu models.TodoUseCaseInterface) {
	handler := &todoHandler{
		todoUseCase: tu,
	}

	r.GET("/", handler.FindAll)
	r.GET("/:id", handler.FindById)
	r.POST("/", handler.Create)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
}

func (h *todoHandler) FindAll(c *gin.Context) {
	todos, err := h.todoUseCase.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	todoResponses := []api.TodoResponse{}
	for _, todo := range todos {
		todoResponses = append(todoResponses, api.TodoResponse{
			ID:   todo.ID,
			Name: todo.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todoResponses,
	})
}

func (h *todoHandler) FindById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID invalid",
		})
		return
	}

	todo, err := h.todoUseCase.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": api.TodoResponse{
			ID:   todo.ID,
			Name: todo.Name,
		},
	})
}

func (h *todoHandler) Create(c *gin.Context) {
	var req api.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo, err := h.todoUseCase.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": api.TodoResponse{
			ID:   todo.ID,
			Name: todo.Name,
		},
	})
}

func (h *todoHandler) Update(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID invalid",
		})
		return
	}

	var req api.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo, err := h.todoUseCase.Update(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": api.TodoResponse{
			ID:   todo.ID,
			Name: todo.Name,
		},
	})
}

func (h *todoHandler) Delete(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID invalid",
		})
		return
	}

	err = h.todoUseCase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
