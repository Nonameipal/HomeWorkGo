package controller

import (
	"fmt"
	"go-shop/internal/models"
)

type Controller struct {
	Categories []models.Category
}


func NewController() *Controller {
	return &Controller{
		Categories: []models.Category{
			{ID: 1, Name: "Продукты", Items: []string{}},
			{ID: 2, Name: "Бытовая химия", Items: []string{}},
			{ID: 3, Name: "Электроника", Items: []string{}},
		},
	}
}


func (c *Controller) PrintCategory(categoryID int) {
	for _, cat := range c.Categories {
		if cat.ID == categoryID {
			fmt.Printf("\n--- Категория: %s ---\n", cat.Name)
			if len(cat.Items) == 0 {
				fmt.Println("Список пуст.")
			} else {
				for i, item := range cat.Items {
					fmt.Printf("%d. %s\n", i+1, item)
				}
			}
			fmt.Println("-----------------------")
			return
		}
	}
	fmt.Println("Ошибка: Категория с таким номером не найдена.")
}


func (c *Controller) AddItem(categoryID int, item string) {
	for i, cat := range c.Categories {
		if cat.ID == categoryID {
			c.Categories[i].Items = append(c.Categories[i].Items, item)
			fmt.Printf("Успешно добавлено '%s' в категорию '%s'\n", item, cat.Name)
			return
		}
	}
	fmt.Println("Ошибка: Категория с таким номером не найдена.")
}




func (c *Controller) showCategories() {
	fmt.Println("\nДоступные категории:")
	for _, cat := range c.Categories {
		fmt.Printf("%d. %s\n", cat.ID, cat.Name)
	}
}