package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		for _, i := range arg {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
