package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	size := max - min
	param := make([]int, size)
	for i := 0; i < size; i++ {
		param[i] = min + i
	}
	return param
}
