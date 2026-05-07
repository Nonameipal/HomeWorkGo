package controller

import (
	"fmt"

	"Abramyan/internal/service"
	"Abramyan/tasks"
)

func GetAllTasks() {
	tasksList := service.GetAllTasks()
	fmt.Println("====== Список задач ======")
	for _, t := range tasksList {
		fmt.Printf("ID: %d | %s - %s\n", t.ID, t.Name, t.Description)
	}
	fmt.Println("==========================")
}

func GetTaskByID() {
	fmt.Print("Введите номер задачи: ")
	var id int
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Ошибка ввода, введите число.")
		return
	}

	task, found := service.GetTaskByID(id)
	if !found {
		fmt.Println("Задача с таким номером не найдена.")
		return
	}

	fmt.Printf("Выбрана задача: %s - %s\n", task.Name, task.Description)
	RunTaskLogic(id)
}

func RunTaskLogic(id int) {
	switch id {
	case 1:
		var a float64
		fmt.Print("Введите a: ")
		fmt.Scanln(&a)
		fmt.Printf("Периметр: %.2f\n", tasks.GetPerimeter(a))
	case 2:
		var a float64
		fmt.Print("Введите a: ")
		fmt.Scanln(&a)
		fmt.Printf("Площадь: %.2f\n", tasks.GetArea(a))
	case 3:
		var a, b float64
		fmt.Print("Введите a и b (через пробел/ентер): ")
		fmt.Scanln(&a, &b)
		area, perim := tasks.GetAreaPerimeterRectangle(a, b)
		fmt.Printf("Площадь: %.2f, Периметр: %.2f\n", area, perim)
	case 4:
		var d float64
		fmt.Print("Введите d: ")
		fmt.Scanln(&d)
		fmt.Printf("Длина окружности: %.2f\n", tasks.GetCircleDiameter(d))
	case 5:
		var a float64
		fmt.Print("Введите a: ")
		fmt.Scanln(&a)
		v, s := tasks.GetCube(a)
		fmt.Printf("Объем: %.2f, Площадь поверхности: %.2f\n", v, s)
	case 6:
		var a, b, c float64
		fmt.Print("Введите a, b c: ")
		fmt.Scanln(&a, &b, &c)
		v, s := tasks.GetBox(a, b, c)
		fmt.Printf("Объем: %.2f, Площадь поверхности: %.2f\n", v, s)
	case 7:
		var r float64
		fmt.Print("Введите R: ")
		fmt.Scanln(&r)
		l, s := tasks.GetCircleRadius(r)
		fmt.Printf("Длина: %.2f, Площадь: %.2f\n", l, s)
	case 8:
		var a, b float64
		fmt.Print("Введите a и b: ")
		fmt.Scanln(&a, &b)
		fmt.Printf("Среднее арифметическое: %.2f\n", tasks.GetArithmeticMean(a, b))
	case 9:
		var a, b float64
		fmt.Print("Введите a и b: ")
		fmt.Scanln(&a, &b)
		fmt.Printf("Среднее геометрическое: %.2f\n", tasks.GetGeometricMean(a, b))
	case 10:
		var a, b float64
		fmt.Print("Введите a и b: ")
		fmt.Scanln(&a, &b)
		sum, diff, mult, div := tasks.GetSquareO(a, b)
		fmt.Printf("Сумма квадратов: %.2f, Разность квадратов: %.2f, Произведение квадратов: %.2f, Частное квадратов: %.2f\n", sum, diff, mult, div)
	case 11:
		var a, b float64
		fmt.Print("Введите a и b: ")
		fmt.Scan(&a, &b)
		sum, diff, mult, div := tasks.GetModule(a, b)
		fmt.Printf("Сумма %.2f, Разность %.2f, произведение %.2f, частное их модулей %.2f\n", sum, diff, mult, div)
	case 12:
		var a, b float64
		fmt.Print("Введите a и b: ")
		fmt.Scan(&a, &b)
		c, p := tasks.GetTriangle(a, b)
		fmt.Printf("Гипотенуза: %.2f, Периметр: %.2f\n", c, p)
	case 13:
		var a, b float64
		fmt.Print("Введите a и b: ")
		fmt.Scan(&a, &b)
		a1, b1, c1 := tasks.GetRingArea(a, b)
		fmt.Printf("Площадь S1: %.2f, Площадь S2: %.2f, Площадь S3: %.2f\n", a1, b1, c1)
	case 14:
		var a float64
		fmt.Print("Введите L: ")
		fmt.Scan(&a)
		r, s := tasks.GetCircleL(a)
		fmt.Printf("Радиус круга: %.2f, Площадь круга: %.2f\n", r, s)
	case 15:
		var a float64
		fmt.Print("Введите площадь круга S: ")
		fmt.Scan(&a)
		d, l := tasks.GetCircleS(a)
		fmt.Printf("Диаметр: %.2f, Длина: %.2f\n", d, l)
	case 16:
		var x1, x2 float64
		fmt.Print("Введите x1 и x2: ")
		fmt.Scan(&x1, &x2)
		d := tasks.GetDistance(x1, x2)
		fmt.Printf("Расстояние: %.2f\n", d)
	case 17:
		var a, b, c float64
		fmt.Print("Введите A, B, C: ")
		fmt.Scan(&a, &b, &c)
		ac, bc, sum := tasks.GetSegmentL(a, b, c)
		fmt.Printf("AC: %.2f, BC: %.2f, Сумма: %.2f\n", ac, bc, sum)
	case 18:
		var a, b, c float64
		fmt.Print("Введите A, B, C: ")
		fmt.Scan(&a, &b, &c)
		prod := tasks.GetPointsP(a, b, c)
		fmt.Printf("Произведение AC и BC: %.2f\n", prod)
	case 19:
		var x1, y1, x2, y2 float64
		fmt.Print("Введите x1, y1, x2, y2: ")
		fmt.Scan(&x1, &y1, &x2, &y2)
		p, s := tasks.GetRectanglePS(x1, y1, x2, y2)
		fmt.Printf("Периметр: %.2f, Площадь: %.2f\n", p, s)
	case 20:
		var x1, y1, x2, y2 float64
		fmt.Print("Введите x1, y1, x2, y2: ")
		fmt.Scan(&x1, &y1, &x2, &y2)
		d := tasks.GetDistance2D(x1, y1, x2, y2)
		fmt.Printf("Расстояние: %.2f\n", d)
	default:
		fmt.Println("Логика для данной задачи еще не реализована.")
	}
}
