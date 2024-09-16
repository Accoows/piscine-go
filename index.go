package piscine

func Index(s string, toFind string) int {
	if len(s) == 0 {
		return -1
	}
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			if len(toFind) == 1 {
				return i
			}
			if len(s[i:]) >= len(toFind) {
				if s[i:i+len(toFind)] == toFind {
					return i
				}
			}
		}
	}
	return -1
}
