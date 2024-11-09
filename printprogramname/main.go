package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := os.Args[0]
	if len(i) > 2 && i[:2] == "./" {
		for _, r := range i[2:] {
			z01.PrintRune(r)
		}
	} else {
		for _, j := range i {
			z01.PrintRune(j)
		}
	}
	z01.PrintRune('\n')
}
