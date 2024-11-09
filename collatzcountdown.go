package piscine

func CollatzCountdown(start int) int {
	i := 0
	if start <= 0 {
		return -1
	}
	for start > 1 {
		if start%2 == 0 {
			start /= 2
			i++
		} else if start%2 != 0 {
			start = (start * 3) + 1
			i++
		}
	}
	return i
}
