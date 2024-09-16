package piscine

func Capitalize(s string) string {
	res_runes := []rune(s)
	if len(s) == 0 {
		return ""
	} else {
		if s[0] >= 'a' && s[0] <= 'z' {
			res_runes[0] = rune(s[0] - 32)
		}
		for i := 0; i <= len(s)-2; i++ {
			if (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') && (s[i] < '0' || s[i] > '9') {
				if res_runes[i+1] >= 'a' && res_runes[i+1] <= 'z' {
					res_runes[i+1] = res_runes[i+1] - 32
				}
			}
			if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
				if res_runes[i+1] >= 'A' && res_runes[i+1] <= 'Z' {
					res_runes[i+1] = res_runes[i+1] + 32
				}
			}
		}
		return string(res_runes)
	}
}
