package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	printargs := os.Args[1:]
	argsindex := len(printargs) - 1
	for argsindex >= 0 {
		for _, mot := range printargs[argsindex] {
			z01.PrintRune(mot)
		}
		argsindex--
		z01.PrintRune('\n')
	}
}
