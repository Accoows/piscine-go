package piscine

func StrRev(s string) string {
	printrev := ""
	for _, i := range s {
		printrev = string(i) + printrev
	}
	return string(printrev)
}
