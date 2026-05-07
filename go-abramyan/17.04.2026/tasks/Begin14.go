package tasks
import "math"
func GetCircleL(l float64) (float64, float64) {
	const pi = 3.14
	r := l / (2 * pi)
	s := pi * math.Pow(r, 2)
	return r, s
}