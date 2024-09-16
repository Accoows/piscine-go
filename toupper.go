package piscine

func ToUpper(s string) string {
	var restoupper string
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			restoupper += string(i - 32)
		} else {
			restoupper += string(i)
		}
	}
	return restoupper
}
