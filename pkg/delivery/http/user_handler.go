package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kurniacf/stunting-be/pkg/api"
	"github.com/kurniacf/stunting-be/pkg/middleware"
	"github.com/kurniacf/stunting-be/pkg/models"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserUsecase models.UserUsecase
}

func NewUserHandler(r *gin.RouterGroup, uu models.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: uu,
	}

	authorized := r.Group("/user", middleware.JwtAuthMiddleware())
	{
		authorized.GET("/", handler.GetUser)
		authorized.PUT("/:id", handler.UpdateUser)
		authorized.DELETE("/:id", handler.DeleteUser)
		// Add more routes as necessary
	}
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	email, _ := c.Get("user_id")

	user, err := uh.UserUsecase.GetByID(uint(email.(float64)))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": api.GetUserResponse{
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uint(id)
	if err := uh.UserUsecase.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := uh.UserUsecase.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "User deleted successfully"})
}
