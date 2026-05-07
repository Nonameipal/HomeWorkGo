package controller

import "fmt"

type Controller struct {
	Tasks    []string
	WorkDays []string
	Weekend  []string
}

func NewController(tasks []string) *Controller {
	return &Controller{
		Tasks:    tasks,
		WorkDays: tasks[0:5],
		Weekend:  tasks[5:7],
	}
}

func (c *Controller) PrintAllTasks() {
	fmt.Println("\n--- Все задачи на неделю ---")
	for _, task := range c.Tasks {
		fmt.Println(task)
	}
}

func (c *Controller) PrintWorkdayTasks() {
	fmt.Println("\n--- Задачи на будни ---")
	for _, task := range c.WorkDays {
		fmt.Println(task)
	}
}

func (c *Controller) PrintWeekendTasks() {
	fmt.Println("\n--- Задачи на выходные ---")
	for _, task := range c.Weekend {
		fmt.Println(task)
	}
}

func (c *Controller) AddWorkdayTask(dayIndex int, newTask string) {
	if dayIndex >= 0 && dayIndex < len(c.WorkDays) {
		c.WorkDays[dayIndex] += ", " + newTask
		fmt.Printf("Успешно добавлено '%s' в будний день!\n", newTask)
	} else {
		fmt.Println("неверный индекс буднего дня.")
	}
}

func (c *Controller) AddWeekendTask(dayIndex int, newTask string) {
	if dayIndex >= 0 && dayIndex < len(c.Weekend) {
		c.Weekend[dayIndex] += ", " + newTask
		fmt.Printf("Успешно добавлено '%s' в выходной день!\n", newTask)
	} else {
		fmt.Println("неверный индекс выходного дня.")
	}
}