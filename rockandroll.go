package piscine

func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative\n"
	}
	var result string
	if n%2 == 0 {
		result += "rock\n"
	}
	if n%3 == 0 {
		result += "roll\n"
	}
	if n%2 == 0 && n%3 == 0 {
		result = "rock and roll\n"
	}
	if result == "" {
		return "error: non divisible\n"
	}
	return result
}
