package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

const a = "x = 42, y = 21"

func str(s string) {
	str := []rune(s)
	for i := 0; i < 14; i++ {
		z01.PrintRune(str[i])
	}
}

func main() {
	points := &point{}
	str(a)
	z01.PrintRune('\n')
	setPoint(points)
}
