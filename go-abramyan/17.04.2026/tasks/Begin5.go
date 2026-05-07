package tasks
import "math"
func GetCube(a float64) (float64, float64) {
	s := math.Pow(a, 3)
	p := 6 * math.Pow(a, 2)
	return s, p
}