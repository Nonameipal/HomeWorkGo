package tasks
import "math"
func GetModule(a, b float64) (float64,  float64, float64, float64){
	absA, absB := math.Abs(a), math.Abs(b)
	return absA + absB, absA - absB, absA * absB, absA / absB
} 