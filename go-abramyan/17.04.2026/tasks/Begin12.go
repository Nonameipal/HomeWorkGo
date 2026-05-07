package tasks

import "math"

func GetTriangle(a, b float64) (float64, float64) {
	c := math.Sqrt(a*a + b*b)
	p := a + b + c
	return c, p
}
