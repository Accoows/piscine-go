package main

import "github.com/01-edu/z01"

func main() {
	var alphabet string = "abcdefghijklmnopqrstuvwxyz"
	for compt := 0; compt < 26; compt++ {
		z01.PrintRune(rune(alphabet[compt]))
	}
	z01.PrintRune('\n')
}
