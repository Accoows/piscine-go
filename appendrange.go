package piscine

func AppendRange(min, max int) []int {
	var param []int
	if min >= max {
		return nil
	}
	for i := min - 1; i < max-1; i++ {
		param = append(param, i+1)
	}
	return param
}
