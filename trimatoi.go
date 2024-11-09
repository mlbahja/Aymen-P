package piscine

func TrimAtoi(s string) int {
	i := 0
	j := 0
	k := 1
	for _, l := range s {
		if l == '-' && j == 0 {
			k = -1
		} else if l >= '0' && l <= '9' {
			i *= 10
			i += int(l - '0')
			j++
		}
	}
	return i * k
}
