package piscine

func LastRune(s string) rune {
	printstr := []rune(s)
	return printstr[len(s)-1]
}
