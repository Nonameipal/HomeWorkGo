package tasks

import "math"
func GetRectanglePS(x1, y1, x2, y2 float64) (float64, float64) {
	a := math.Abs(x2 - x1)
	b := math.Abs(y2 - y1)
	p := 2 * (a + b)
	s := a * b
	return p, s
}