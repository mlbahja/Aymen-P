package piscine

func FirstRune(s string) rune {
	if len(s) == 0 {
		return 0
	}

	k := []rune(s)
	return k[0]
}
