package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	var w string
	var sw []string
	for _, char := range str {
		if char != ' ' {
			w += string(char)
		} else {
			sw = append(sw, w)
			w = ""
		}
	}
	sw = append(sw, w)
	summary := make(map[string]int)
	for _, w := range sw {
		summary[w]++
	}
	return summary
}
