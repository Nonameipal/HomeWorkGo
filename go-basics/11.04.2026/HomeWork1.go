package main
import "fmt"
func sum(a, b int) int {
	return a + b
}
func perimeter(c, d float64) float64 {
	return 2 * (c + d)
}
func area(s float64) float64 {	
	return s * s
}
func main() {
	linii := "====================================="

	result := sum(5, 10)
	fmt.Println("Задание 1 (сумма чисел): ", result)
	fmt.Println(linii)

	result2 := perimeter(5, 3)
	fmt.Println("Задание 2 (Периметр прямоугольника): ", result2)
	fmt.Println(linii)

	result3 := area(5)
	fmt.Println("Задание 3 (Площадь квадрата): ", result3)
}