package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Sort(char []string) {
	char_index := len(char) - 1
	for i := 0; i < char_index; i++ {
		for j := i + 1; j < len(char); j++ {
			if char[i] > char[j] {
				char[i], char[j] = char[j], char[i]
			}
		}
	}
}

func main() {
	printargs := os.Args[1:]
	Sort(printargs)
	for _, posargs := range printargs {
		for _, mot := range posargs {
			z01.PrintRune(mot)
		}
		z01.PrintRune('\n')
	}
}
