package piscine

func SplitWhiteSpaces(s string) []string {
	str := ""
	slice := []string{}
	for _, j := range s {
		if j != ' ' {
			str = str + string(j)
		} else if str != "" {
			slice = append(slice, str)
			str = ""
		}
	}
	if str != "" {
		slice = append(slice, str)
	}
	return slice
}
