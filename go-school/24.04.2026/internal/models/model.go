package models

import "fmt"

type StudentGrades struct {
	Grades [5]int
}

// AddGrade добавляет оценку в массив по указанному индексу
// Параметры:
//   - grade: оценка (не должна быть выше 10 баллов)
//   - index: порядковый номер (индекс) в массиве (0-4)
// Возвращает результат: true если оценка добавлена успешно, false если оценка отклонена
func (sg *StudentGrades) AddGrade(grade, index int) bool {
	// Проверяем, что индекс в пределах массива
	if index < 0 || index >= len(sg.Grades) {
		fmt.Printf("Ошибка: индекс %d вне диапазона [0-4]\n", index)
		return false
	}

	// Проверяем, что оценка не выше 10 баллов
	if grade > 10 {
		fmt.Printf("Ошибка: оценка %d выше допустимого максимума (10 баллов)\n", grade)
		return false
	}

	// Проверяем, что оценка неотрицательна
	if grade < 0 {
		fmt.Printf("Ошибка: оценка %d не может быть отрицательной\n", grade)
		return false
	}

	// Добавляем оценку в массив
	sg.Grades[index] = grade
	fmt.Printf("Оценка %d успешно добавлена на позицию %d\n", grade, index)
	return true
}

// GetGrades возвращает все оценки студента
func (sg *StudentGrades) GetGrades() [5]int {
	return sg.Grades
}

// GetAverageGrade вычисляет и возвращает среднюю оценку
func (sg *StudentGrades) GetAverageGrade() float64 {
	sum := 0
	count := 0
	for _, grade := range sg.Grades {
		if grade > 0 {
			sum += grade
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}
