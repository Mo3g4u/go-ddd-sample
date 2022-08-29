package main

import (
	"log"

	"github.com/Mo3g4u/go-ddd-sample/infrastructure/repository"
	"github.com/Mo3g4u/go-ddd-sample/infrastructure/service"
	"github.com/Mo3g4u/go-ddd-sample/interfaces/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env ファイルの読込の失敗しました。")
	}

	db, err := repository.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	router := gin.Default()
	r := repository.NewRepository(db)
	us := service.NewUserService(r)

	uh := handler.NewUserHandler(r, us)
	router.POST("/users", uh.Create)

	router.Run()

}
