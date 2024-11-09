package piscine

func BasicAtoi(s string) int {
	res := 0
	for _, n := range s {
		res = res*10 + int(n-'0')
	}
	return res
}
