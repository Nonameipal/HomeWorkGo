package controller

import (
	"fmt"
	"os"
	"os/exec"
)

func InitRoutes() {
	fmt.Println("Добро пожаловать в Решатель Задач!")
	for {
		fmt.Println("\n====== Меню ======")
		fmt.Println("0. Выход -_0")
		fmt.Println("1. Получить все задачи")
		fmt.Println("2. Получить задание по номеру")
		fmt.Println("==================")
		
		fmt.Print("Выберите нужный пункт: ")
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "0":
			fmt.Println("Мем) след раз рекрол)")
			cmdExec := exec.Command("curl", "parrot.live")
            cmdExec.Stdout = os.Stdout
            cmdExec.Run()
		case "1":
			GetAllTasks()
		case "2":
			GetTaskByID()
		default:
			fmt.Println("Несуществующая команда, попробуйте еще раз")
		}
	}
}
