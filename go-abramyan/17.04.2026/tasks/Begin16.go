package tasks

import "math"

func GetDistance(x1, x2 float64) float64 {
	return math.Abs(x2 - x1)
}
