package controller
import (
	"fmt"
)

func (c *Controller) Run() {
	for {
		fmt.Println("\n==== СПИСОК ПОКУПОК =====")
		fmt.Println("1. Посмотреть список покупок")
		fmt.Println("2. Добавить покупку")
		fmt.Println("3. Выход")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			c.showCategories()
			fmt.Print("Введите порядковый номер категории для просмотра: ")
			var catID int
			fmt.Scan(&catID)
			c.PrintCategory(catID)
		case 2:
			c.showCategories()
			fmt.Print("Введите порядковый номер категории для добавления: ")
			var catID int
			fmt.Scan(&catID)
			fmt.Print("Введите название покупки: ")
			var itemName string
			fmt.Scan(&itemName)
			c.AddItem(catID, itemName)
		case 3:
			fmt.Println("Выход из программы. ")
			return
		default:
			fmt.Println("Неизвестная команда.ё")
		}
	}
}
