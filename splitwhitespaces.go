package piscine

func SplitWhiteSpaces(s string) []string {
	var res []string
	var mot string
	for i, charact := range s {
		if charact != ' ' && charact != '\t' && charact != '\n' {
			mot += string(charact)
		}
		if ((charact == ' ' || charact == '\t' || charact == '\n') && mot != "") || i == len(s)-1 {
			res = append(res, mot)
			mot = ""
		}
	}
	return res
}
