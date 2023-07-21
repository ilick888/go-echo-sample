package main

import (
	"go-sample/controller"
	"go-sample/db"
	"go-sample/repository"
	"go-sample/router"
	"go-sample/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)

	e.Logger.Debug(e.Start(":8080"))

}
