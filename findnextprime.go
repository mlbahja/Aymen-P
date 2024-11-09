package piscine

func FindNextPrime(nb int) int {
	var i int
	if nb < 0 {
		i = 2
	} else {
		i = nb
	}

	for !IsPrime(i) {
		i++
	}

	return i
}
