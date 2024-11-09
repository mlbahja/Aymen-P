package piscine

func JumpOver(str string) string {
	s := ""
	for k := 1; k <= len(str); k++ {
		if k%3 == 0 {
			s = s + string(str[k-1])
		}
	}
	s = s + "\n"
	return s
}
