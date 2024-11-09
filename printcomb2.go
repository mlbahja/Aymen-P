package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := ('0' - 48); a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			z01.PrintRune((a / 10) + '0')
			z01.PrintRune((a % 10) + '0')
			z01.PrintRune(' ')
			z01.PrintRune((b / 10) + '0')
			z01.PrintRune((b % 10) + '0')
			if a != 98 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
