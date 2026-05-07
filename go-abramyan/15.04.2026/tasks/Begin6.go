package tasks

func GetBox(a, b, c float64) (float64, float64) {
	s := a * b * c
	p := 2 * (a*b + a*c + b*c)
	return s, p
}