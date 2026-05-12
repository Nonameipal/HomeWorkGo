package main

import (
	"fmt"
	"time"

	"go-todo/internal/task"
)

func main() {
	list := task.NewTaskList()


	t1 := task.New("Купить продукты", "Alice")
	t2 := task.New("Сделать домашнее задание", "Bob")
	t3 := task.New("Позвонить врачу", "Alice")


	t1.Complete()
	t2.SetDeadline(time.Now().Add(24 * time.Hour))
	t3.SetDescription("Позвонить врачу и записаться на приём")

	list.Add(t1)
	list.Add(t2)
	list.Add(t3)

	fmt.Printf("Всего задач в хранилище: %d\n\n", list.Count())


	fmt.Println("=== Все задачи ===")
	for _, t := range list.GetAll() {
		fmt.Printf("  [%d] %s | Выполнена: %v | Автор: %s\n",
			t.ID, t.Description, t.Completed, t.CreatedBy)
	}


	fmt.Println("\n=== Получить задачу по ID=2 ===")
	if found, ok := list.GetByID(2); ok {
		fmt.Printf("  Найдена: %s (дедлайн: %s)\n",
			found.Description, found.Deadline.Format("02.01.2006 15:04"))
	}


	t2.SetDescription("Сделать домашнее задание по Go")
	updated := list.Update(t2)
	fmt.Printf("\n=== Обновление задачи ID=2: успешно=%v ===\n", updated)


	deleted := list.DeleteByID(1)
	fmt.Printf("\n=== Удаление задачи ID=1: успешно=%v ===\n", deleted)
	fmt.Printf("Задач после удаления: %d\n", list.Count())
}
