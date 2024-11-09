package piscine

func DivMod(a, b int, div, mod *int) {
	if b != 0 {
		*div = a / b
		*mod = a % b
	} else {
		*div = 0
		*mod = 0
	}
}
