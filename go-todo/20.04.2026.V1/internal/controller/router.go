package controller

import (
	"fmt"
	"os"
)

func InitRouter(c *TaskController) {
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
			c.GetAllTasks()
		case "2":
			c.AddTask()
		case "3":
			c.DeleteTask()
		case "4":
			c.CompleteTask()
		default:
			fmt.Println("Неверная команда")
		}
	}
}
