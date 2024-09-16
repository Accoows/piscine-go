package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	if nb == 2 {
		return 0
	}
	if nb > 0 {
		printsqrt := 0
		res := 1
		for i := 1; res <= nb; i++ {
			res = i * i
			printsqrt++
			if res == nb {
				return printsqrt
			}
		}
	}
	return 0
}
