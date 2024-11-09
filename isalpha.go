package piscine

func IsAlpha(s string) bool {
	for _, d := range s {
		if (d < 'A' || d > 'Z') && (d < 'a' || d > 'z') && (d < '0' || d > '9') {
			return false
		}
	}
	return true
}
