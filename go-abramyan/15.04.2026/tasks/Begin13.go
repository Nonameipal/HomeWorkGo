package tasks

import "math"
func GetRingArea(a, b float64) (float64, float64, float64) {
	const pi = 3.14
	a1 := pi * math.Pow(a, 2)
    b1 := pi * math.Pow(b, 2)
	return a1, b1, a1 - b1
}