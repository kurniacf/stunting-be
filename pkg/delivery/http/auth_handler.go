package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kurniacf/stunting-be/pkg/api"
	"github.com/kurniacf/stunting-be/pkg/models"
	"github.com/kurniacf/stunting-be/pkg/usecase"
)

type AuthHandler struct {
	AuthUsecase usecase.AuthUsecase
	UserUsecase models.UserUsecase
}

func NewAuthHandler(r *gin.RouterGroup, au usecase.AuthUsecase, uu models.UserUsecase) {
	handler := &AuthHandler{
		AuthUsecase: au,
		UserUsecase: uu,
	}

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var req api.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := ah.UserUsecase.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var req api.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ah.AuthUsecase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": api.LoginResponse{
			Token: token,
		},
	})
}
