package piscine

func ToLower(s string) string {
	var restolower string
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			restolower += string(i + 32)
		} else {
			restolower += string(i)
		}
	}
	return restolower
}
