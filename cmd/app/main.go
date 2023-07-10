package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kurniacf/stunting-be/configs"
	"github.com/kurniacf/stunting-be/pkg/delivery/http"
	"github.com/kurniacf/stunting-be/pkg/repository"
	"github.com/kurniacf/stunting-be/pkg/usecase"
)

func main() {
	db := configs.InitDB()

	userRepo := repository.NewMysqlUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	r := gin.Default()

	v1 := r.Group("/v1")
	http.NewUserHandler(v1, userUsecase)

	r.Run()
}
