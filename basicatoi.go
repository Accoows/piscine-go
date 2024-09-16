package piscine

func BasicAtoi(s string) int {
	num := 0
	for _, r := range s {
		num = num*10 + int(r-'0')
	}
	return num
}
