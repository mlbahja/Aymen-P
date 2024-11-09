package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	maxValue := a[0]
	for _, value := range a {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}
