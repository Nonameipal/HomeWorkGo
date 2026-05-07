package controller

import (
	"Todo/internal/models"
	"fmt"
	"os"
)

func InitTasks() {
	addTask("Сделать домашнее задание")
	addTask("Пойти в магазин")
	addTask("Курс по Go")
}

func addTask(desc string) {
	task := &models.Task{
		Id:          models.NextID,
		Description: desc,
		Completed:   false,
	}
	models.Tasks[models.NextID] = task
	models.NextID++
}

func InitRouter() {
	showMenu()
}

func showMenu() {
	for {
		fmt.Println("=========Список задач=========")
		fmt.Println("1. Просмотреть задачи")
		fmt.Println("2. Добавить задачу")
		fmt.Println("3. Удалить задачу")
		fmt.Println("4. Задача выполнена id")
		fmt.Println("0. Выход")
		fmt.Println("================================")

		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "0":
			os.Exit(0)
		case "1":
			GetAllTasks()
		case "2":
			AddTask()
		case "3":
			DeleteTask()
		case "4":
			CompleteTask()
		default:
			fmt.Println("Неверная команда")
		}
	}
}
