package piscine

func IsLower(s string) bool {
	for _, d := range s {
		if d < 'a' || d > 'z' {
			return false
		}
	}
	return true
}
