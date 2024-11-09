package piscine

func BasicAtoi2(s string) int {
	res := 0
	for _, n := range s {
		if n >= '0' && n <= '9' && n != 32 {
			res = res*10 + int(n-'0')
		} else {
			return 0
		}
	}
	return res
}
