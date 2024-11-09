package piscine

func IsNumeric(s string) bool {
	for _, h := range s {
		if h < '0' || h > '9' {
			return false
		}
	}
	return true
}
