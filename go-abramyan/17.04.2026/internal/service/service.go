package service

import (
	"Abramyan/internal/models"
)

var tasksList = []models.Task{
	{ID: 1, Name: "Begin1", Description: "Найти периметр квадрата P = 4*a"},
	{ID: 2, Name: "Begin2", Description: "Найти площадь квадрата S = a^2"},
	{ID: 3, Name: "Begin3", Description: "Найти площадь и периметр прямоугольника: S = a*b, P = 2*(a+b)"},
	{ID: 4, Name: "Begin4", Description: "Найти длину окружности L = pi*d"},
	{ID: 5, Name: "Begin5", Description: "Найти объем и площадь поверхности куба"},
	{ID: 6, Name: "Begin6", Description: "Найти объем и площадь поверхности параллелепипеда"},
	{ID: 7, Name: "Begin7", Description: "Найти длину и площадь круга"},
	{ID: 8, Name: "Begin8", Description: "Найти среднее арифметическое"},
	{ID: 9, Name: "Begin9", Description: "Найти среднее геометрическое"},
	{ID: 10, Name: "Begin10", Description: "Сумма, разность, произведение и частное квадратов двух чисел"},
	{ID: 11, Name: "Begin11", Description: "Сумма, разность, произведение и частное модулей двух чисел"},
	{ID: 12, Name: "Begin12", Description: "Гипотенуза и периметр прямоугольного треугольника"},
	{ID: 13, Name: "Begin13", Description: "Площади двух кругов и площадь кольца"},
	{ID: 14, Name: "Begin14", Description: "Радиус и площадь по длине окружности"},
	{ID: 15, Name: "Begin15", Description: "Диаметр и длина окружности по площади круга"},
	{ID: 16, Name: "Begin16", Description: "Расстояние между двумя точками на числовой оси"},
	{ID: 17, Name: "Begin17", Description: "Длины отрезков AC и BC и их сумма"},
	{ID: 18, Name: "Begin18", Description: "Произведение длин отрезков AC и BC"},
	{ID: 19, Name: "Begin19", Description: "Периметр и площадь прямоугольника по координатам"},
	{ID: 20, Name: "Begin20", Description: "Расстояние между двумя точками на плоскости"},
}

func GetAllTasks() []models.Task {
	return tasksList
}

func GetTaskByID(id int) (models.Task, bool) {
	for _, t := range tasksList {
		if t.ID == id {
			return t, true
		}
	}
	return models.Task{}, false
}
