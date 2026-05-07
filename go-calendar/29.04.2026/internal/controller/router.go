package controller

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (c *Controller) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n==== МИНИ-КАЛЕНДАРЬ =====")
		fmt.Println("1. Прочитать задачи на всю неделю")
		fmt.Println("2. Прочитать задачи на будни")
		fmt.Println("3. Прочитать задачи на выходные")
		fmt.Println("4. Добавить задачу на будний день")
		fmt.Println("5. Добавить задачу на выходной день")
		fmt.Println("6. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			c.PrintAllTasks()
		case 2:
			c.PrintWorkdayTasks()
		case 3:
			c.PrintWeekendTasks()
		case 4:
			fmt.Print("Введите индекс буднего дня (0 - Понедельник, 4 - Пятница): ")
			var dayIndex int
			fmt.Scan(&dayIndex)

			fmt.Print("Введите задачу: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			
			c.AddWorkdayTask(dayIndex, task)
		case 5:
			fmt.Print("Введите индекс выходного дня (0 - Суббота, 1 - Воскресенье): ")
			var dayIndex int
			fmt.Scan(&dayIndex)

			fmt.Print("Введите задачу: ")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)

			c.AddWeekendTask(dayIndex, task)
		case 6:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}