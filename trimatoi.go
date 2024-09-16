package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0

	for _, charactere := range s {
		if charactere == '-' && result == 0 {
			sign = -1
		}
		if charactere >= '0' && charactere <= '9' {
			result = result*10 + int(charactere-'0') // On construit l'entier à partir des chiffres
		}
	}
	return result * sign
}
