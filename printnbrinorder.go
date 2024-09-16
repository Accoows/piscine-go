package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}

	var resdigits []int
	for n > 0 {
		resdigits = append(resdigits, n%10)
		n /= 10
	}

	for i := 0; i < len(resdigits); i++ {
		for j := 0; j < len(resdigits)-i-1; j++ {
			if resdigits[j] > resdigits[j+1] {
				resdigits[j], resdigits[j+1] = resdigits[j+1], resdigits[j]
			}
		}
	}

	for _, printdigits := range resdigits {
		z01.PrintRune(rune(printdigits + '0'))
	}
}
