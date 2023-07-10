package http

import (
	"github.com/gin-gonic/gin"
	"github.com/kurniacf/stunting-be/pkg/models"
)

type UserHandler struct {
	UserUsecase models.UserUsecase
}

func NewUserHandler(r *gin.RouterGroup, uu models.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: uu,
	}
	r.GET("/users", handler.FindUsers)
	// Here you can add more routes for user
}

func (uh *UserHandler) FindUsers(c *gin.Context) {
	// Implement your handler here
}
