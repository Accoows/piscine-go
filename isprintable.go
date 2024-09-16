package piscine

func IsPrintable(s string) bool {
	for _, i := range s {
		if i >= ' ' && i <= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}
