package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	printprog := []rune(os.Args[0])
	res_prog := 0
	for i := len(printprog) - 1; i >= 0; i-- {
		if printprog[i] == '\\' || printprog[i] == '/' {
			res_prog = i + 1
			break
		}
	}
	for i := res_prog; i < len(printprog); i++ {
		if printprog[i] == '.' {
			break
		}
		z01.PrintRune(printprog[i])
	}
	z01.PrintRune('\n')
}
