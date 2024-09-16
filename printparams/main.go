package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	printargs := os.Args[1:]
	for _, posargs := range printargs {
		for _, mot := range posargs {
			z01.PrintRune(mot)
		}
		z01.PrintRune('\n')
	}
}
