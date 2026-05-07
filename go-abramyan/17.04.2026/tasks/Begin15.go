package tasks

import "math"
func GetCircleS(s float64) (float64, float64) {
	const pi = 3.14
	r := math.Sqrt(s / pi)
	d := 2 * r
	l := 2 * pi * r
	return d, l
}