package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		for _, int64 := range "-9223372036854775808" {
			z01.PrintRune(int64)
		}
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	printNumber(n)
}

func printNumber(n int) {
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}
	for j := len(digits) - 1; j >= 0; j-- {
		z01.PrintRune(rune(digits[j] + '0'))
	}
}
