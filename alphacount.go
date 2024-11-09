package piscine

func AlphaCount(s string) int {
	str := []rune(s)
	j := 0
	for i := 0; i < len(s); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			j++
		}
	}
	return j
}
