package piscine

func ToUpper(s string) string {
	d := []rune(s)
	for i, j := range s {
		if j >= 'a' && j <= 'z' {
			d[i] = j - ('a' - 'A')
		}
	}
	return string(d)
}
