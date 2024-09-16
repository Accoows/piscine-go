package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func nbrRune(nbrpoint int, strnb string) string {
	len_nb := len(strnb)
	if nbrpoint == 0 {
		return "0"
	}
	res := ""
	for nbrpoint > 0 {
		res = string(strnb[nbrpoint%len_nb]) + res
		nbrpoint = nbrpoint / len_nb
	}
	return res
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	res := "x = " + nbrRune(points.x, "0123456789") + ", y = " + nbrRune(points.y, "0123456789")

	for _, i := range res {
		z01.PrintRune(i)
	}
}
