package piscine

func NRune(s string, n int) rune {
	k := []rune(s)
	if n <= 0 || n > len(k) {
		return 0
	}
	return k[n-1]
}
