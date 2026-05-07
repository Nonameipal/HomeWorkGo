package controller

import (
	"Todo/internal/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetAllTasks() {
	if len(models.Tasks) == 0 {
		fmt.Println("Задачи нет")
		return
	}
	for _, task := range models.Tasks {
		status := "Не выполнено"
		if task.Completed {
			status = "Выполнено"
		}
		fmt.Printf("%d. [%s] %s\n", task.Id, status, task.Description)
	}
}

func AddTask() {
	fmt.Print("Введите описание задачи: ")
	reader := bufio.NewReader(os.Stdin)
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	if description == "" {
		fmt.Println("Описание не может быть пустым")
		return
	}
	task := &models.Task{
		Id:          models.NextID,
		Description: description,
		Completed:   false,
	}
	models.Tasks[models.NextID] = task
	fmt.Println("Задача добавлена, ID:", models.NextID)
	models.NextID++
}

func DeleteTask() {
	fmt.Print("Введите id задачи для удаления: ")
	var id int
	fmt.Scanln(&id)
	if _, exists := models.Tasks[id]; exists {
		delete(models.Tasks, id)
		fmt.Println("Задача удалена")
	} else {
		fmt.Println("Задача не найдена")
	}
}

func CompleteTask() {
	fmt.Print("Введите id задачи для отметки как выполненная: ")
	var id int
	fmt.Scanln(&id)
	if task, exists := models.Tasks[id]; exists {
		task.Completed = true
		fmt.Println("Задача отмечена как выполненная")
	} else {
		fmt.Println("Задача не найдена")
	}
}
