package tasks
func GetSquareO(a, b float64) (float64, float64, float64, float64) {
	sqA := a * a
	sqB := b * b
	
	sum := sqA + sqB
	diff := sqA - sqB
	mult := sqA * sqB
	div := sqA / sqB
	
	return sum, diff, mult, div
}