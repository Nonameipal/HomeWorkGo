package app

import (
	"Todo/internal/controller"
	"Todo/internal/repository"
	"Todo/internal/service"
)

func Run() {
	repo := repository.NewTaskRepository()

	svc := service.NewTaskService(repo)

	ctrl := controller.NewTaskController(svc)

	controller.InitRouter(ctrl)
}
