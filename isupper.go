package piscine

func IsUpper(s string) bool {
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
