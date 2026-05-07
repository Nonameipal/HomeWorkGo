package tasks

import "math"

func GetSegmentL(a, b, c float64) (float64, float64, float64) {
	ac := math.Abs(c - a)
	bc := math.Abs(c - b)
	sum := ac + bc
	return ac, bc, sum
}