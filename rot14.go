package piscine

func Rot14(s string) string {
	str := ""
	for _, i := range s {
		if (i >= 'a' && i <= 'l') || (i >= 'A' && i <= 'L') {
			str += string(i + 14)
		} else if (i >= 'n' && i <= 'z') || (i >= 'N' && i <= 'Z') {
			str += string(i - 12)
		} else if i == 'M' {
			str += string('A')
		} else if i == 'm' {
			str += string('a')
		} else {
			str += string(i)
		}
	}
	return str
}
