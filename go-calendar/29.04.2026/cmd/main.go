package main

import (
	"go-calendar/internal/controller"
	"go-calendar/internal/models"
)

func main() {

	tasks := models.DefaultTasks()
	app := controller.NewController(tasks)
	app.Run()
}
