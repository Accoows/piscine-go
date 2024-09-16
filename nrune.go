package piscine

func NRune(s string, n int) rune {
	res_runes := []rune(s)
	if n > len(s) || n <= 0 {
		return 0
	}
	return res_runes[n-1]
}
