package piscine

func ToLower(s string) string {
	d := []rune(s)
	for i, j := range s {
		if j >= 'A' && j <= 'Z' {
			d[i] = j - ('A' - 'a')
		}
	}
	return string(d)
}
