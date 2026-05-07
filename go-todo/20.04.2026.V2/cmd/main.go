package main

import (
	"Todo/internal/controller"
)

func main() {
	controller.InitTasks()
	controller.InitRouter()
}
