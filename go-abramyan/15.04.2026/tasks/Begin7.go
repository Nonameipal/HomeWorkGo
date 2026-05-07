package tasks
import "math"
func GetCircleRadius(a float64) (float64, float64) {
	const pi = 3.14
	l := 2 * pi * a
	s := pi * math.Pow(a, 2)
	return l, s
}