package main

import (
	"go-sample/controller"
	"go-sample/db"
	"go-sample/repository"
	"go-sample/router"
	"go-sample/usecase"
	"go-sample/validator"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	taskRepository := repository.NewTaskRepository(db)
	taskValidator := validator.NewTaskValidator()
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)

	e.Use(middleware.Logger())
	e.Debug = true

	e.Logger.Debug(e.Start(":8080"))

}
