package main

import (
	"fmt"
	"school/internal/models"
)

func main() {
	student := &models.StudentGrades{}
	fmt.Println("=== Система ввода оценок студента ===\n")
	testCases := []struct {
		grade int
		index int
	}{
		{grade: 8, index: 0},
		{grade: 9, index: 1},
		{grade: 12, index: 2},
		{grade: 7, index: 2},
		{grade: 10, index: 3},
		{grade: 6, index: 4},
		{grade: 5, index: 5},
	}

	fmt.Println("Попытки добавления оценок:")
	for _, tc := range testCases {
		student.AddGrade(tc.grade, tc.index)
	}

	fmt.Println("\n=== Результаты ===")
	fmt.Printf("Все оценки: %v\n", student.GetGrades())
	fmt.Printf("Средняя оценка: %.2f\n", student.GetAverageGrade())
}
