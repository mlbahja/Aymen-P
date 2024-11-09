package piscine

func StrRev(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; {
		r[i], r[j] = r[j], r[i]
		j--
		i++
	}
	return string(r)
}
