package piscine

func AlphaCount(s string) int {
	var compteur int
	for _, i := range s {
		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') {
			compteur++
		}
	}
	return compteur
}
