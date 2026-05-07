package controller

import (
	"Todo/internal/service"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TaskController struct {
	service *service.TaskService
}

func NewTaskController(service *service.TaskService) *TaskController {
	return &TaskController{service: service}
}

func (c *TaskController) GetAllTasks() {
	tasklist := c.service.GetAllTasks()
	if len(tasklist) == 0 {
		fmt.Println("Задачи нет")
	} else {
		for _, task := range tasklist {
			status := "Не выполнено"
			if task.Completed {
				status = "Выполнено"
			}
			fmt.Printf("%d. %s %s\n", task.Id, status, task.Description)
		}
	}
}

func (c *TaskController) AddTask() {
	fmt.Println("Введеите описание задачи: ")
	reader := bufio.NewReader(os.Stdin)
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	if description == "" {
		fmt.Println("Описание не может быть пустым")
		return
	}
	id := c.service.AddTask(description)
	fmt.Println("Задача добавлена id задачи: ", id)
}

func (c *TaskController) DeleteTask() {
	fmt.Println("Введите id задачи для удаление: ")
	var id int
	fmt.Scanln(&id)
	if id <= 0 {
		fmt.Println("Неверный id")
		return
	}
	if c.service.DeleteTask(id) {
		fmt.Println("Задача удалена")
	} else {
		fmt.Println("Задача не найдена")
	}
}

func (c *TaskController) CompleteTask() {
	fmt.Println("Введите id задачи для отметки как выполненная: ")
	var id int
	fmt.Scanln(&id)
	if id <= 0 {
		fmt.Println("Неверный id")
		return
	}

	if c.service.CompleteTask(id) {
		fmt.Println("Задача отмечена как выполненная")
	} else {
		fmt.Println("Задача не найдена")
	}
}
