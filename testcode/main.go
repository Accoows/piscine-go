package main

import "piscine"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(piscine.PrintNbr, a)
}

func ForEach(f func(int), a []int) {
	for _, i := range a {
		f(i)
	}
}
