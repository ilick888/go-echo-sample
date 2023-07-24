package main

import (
	"go-sample/controller"
	"go-sample/db"
	"go-sample/repository"
	"go-sample/router"
	"go-sample/usecase"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)

	e.Use(middleware.Logger())
	e.Debug = true

	e.Logger.Debug(e.Start(":8080"))

}
