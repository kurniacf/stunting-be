package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kurniacf/stunting-be/pkg/api"
	"github.com/kurniacf/stunting-be/pkg/models"
)

type childHandler struct {
	ChildUseCase models.ChildUseCaseInterface
}

func NewChildHandler(r *gin.RouterGroup, cu models.ChildUseCaseInterface) {
	handler := &childHandler{
		ChildUseCase: cu,
	}

	r.GET("/user", handler.FindByUserId)
	r.GET("/:id", handler.FindById)
	r.POST("/", handler.Create)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
}

func (h *childHandler) FindByUserId(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token invalid",
		})
		return
	}

	childs, err := h.ChildUseCase.FindByUserId(uint(userId.(float64)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": childs,
	})
}

func (h *childHandler) FindById(c *gin.Context) {
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

	child, err := h.ChildUseCase.FindById(uint(userId.(float64)), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": api.ChildResponse{
			ID:           child.ID,
			Name:         child.Name,
			HealthStatus: child.HealthStatus,
			BirthDate:    child.BirthDate,
			UserId:       child.UserID,
		},
	})
}

func (h *childHandler) Create(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token invalid",
		})
		return
	}

	var req api.CreateChildRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	child, err := h.ChildUseCase.Create(uint(userId.(float64)), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": api.ChildResponse{
			ID:           child.ID,
			Name:         child.Name,
			HealthStatus: child.HealthStatus,
			BirthDate:    child.BirthDate,
			UserId:       child.UserID,
		},
	})
}

func (h *childHandler) Update(c *gin.Context) {
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

	child, err := h.ChildUseCase.FindById(uint(userId.(float64)), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": api.ChildResponse{
			ID:           child.ID,
			Name:         child.Name,
			HealthStatus: child.HealthStatus,
			BirthDate:    child.BirthDate,
			UserId:       child.UserID,
		},
	})
}

func (h *childHandler) Delete(c *gin.Context) {
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

	err = h.ChildUseCase.Delete(uint(userId.(float64)), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
