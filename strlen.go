package piscine

func StrLen(s string) int {
	compt := 0
	for _, i := range s {
		if i == i {
			compt++
		}
	}
	return compt
}
