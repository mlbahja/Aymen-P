package piscine

func IsUpper(s string) bool {
	for _, d := range s {
		if d <= 'A' || d >= 'Z' {
			return false
		}
	}
	return true
}
