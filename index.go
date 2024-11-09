package piscine

func Index(s string, toFind string) int {
	x := []rune(s)
	r := []rune(toFind)
	if len(r) > len(x) {
		return -1
	}
	for i := 0; i < len(x)-len(r); i++ {
		if string(x[i:i+len(r)]) == string(r) {
			return i
		}
	}
	return -1
}
