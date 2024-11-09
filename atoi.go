package piscine

func Atoi(s string) int {
	res := 0
	sign := 1
	for index, value := range s {
		if value == '-' && index == 0 {
			sign = -1
		} else if value == '+' && index == 0 {
			sign = 1
		} else if value >= 'a' && value <= 'z' {
			return 0
		} else if value >= '0' && value <= '9' {
			res *= 10
			res += int(value - '0')
		} else {
			return 0
		}
	}
	return (res * sign)
}
