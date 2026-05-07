package tasks
func GetAreaPerimeterRectangle(a, b float64) (float64, float64) {
	s := a * b
	p := 2 * (a + b)
	return s, p
}