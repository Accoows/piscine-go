package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	taille := 0

	for i := range args {
		taille = i + 1
	}

	if taille > 1 {
		fmt.Println("Too many arguments")
	} else if taille == 0 {
		fmt.Println("File name missing")
	} else if args[0] == "quest8.txt" {
		content, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println("Erreur")
		}
		fmt.Println(string(content))
	}
}
