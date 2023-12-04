package algorithms

func LCM(x, y int) int {
	result := x / GCD(x, y) * y
	return result
}
