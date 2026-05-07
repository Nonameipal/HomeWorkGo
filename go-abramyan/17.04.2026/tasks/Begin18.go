package tasks

import "math"
func GetPointsP(a, b, c float64) float64 {
	ac := math.Abs(c - a)
	bc := math.Abs(c - b)
	return ac * bc
}