package piscine

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}

	k := []rune(s)
	return k[len(s)-1]
}
