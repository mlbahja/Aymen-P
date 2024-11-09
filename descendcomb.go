package piscine

import "github.com/01-edu/z01"

func PrintNumber(n int) {
	res := ""
	if n == 0 {
		res = "0"
	}
	for n != 0 {
		res = string((n%10)+'0') + res
		n /= 10
	}
	if len(res) < 2 {
		z01.PrintRune('0')
	}
	for _, c := range res {
		z01.PrintRune(c)
	}
}

func DescendComb() {
	for a := 99; a >= 1; a-- {
		for b := a - 1; b >= 0; b-- {
			PrintNumber(a)
			z01.PrintRune(' ')
			PrintNumber(b)
			if a != 1 || b != 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
