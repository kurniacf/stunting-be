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
	flag.Parse()

	db := configs.InitDB(*seed)

	userRepo := repository.NewMysqlUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	r := gin.Default()

	api := r.Group("/api")
	httpDeliver.NewUserHandler(api, userUsecase)

	r.Run()
}
