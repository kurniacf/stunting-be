package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/kurniacf/stunting-be/configs"
	httpDeliver "github.com/kurniacf/stunting-be/pkg/delivery/http"
	"github.com/kurniacf/stunting-be/pkg/repository"
	"github.com/kurniacf/stunting-be/pkg/usecase"
)

func main() {
	seed := flag.Bool("seed", false, "Seed the database")
	prod := flag.Bool("prod", false, "Use production database")
	flag.Parse()

	db := configs.InitDB(*seed, *prod)

	userRepo := repository.NewMysqlUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	authUsecase := usecase.NewAuthUsecase(userUsecase)

	r := gin.Default()

	api := r.Group("/api")

	httpDeliver.NewUserHandler(api, userUsecase)

	// Create a new group for auth
	auth := api.Group("/auth")
	httpDeliver.NewAuthHandler(auth, authUsecase, userUsecase)

	r.Run()
}
