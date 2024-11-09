package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	grd := true
	min := true

	for i := 1; i < len(a); i++ {
		if f(a[i], a[i-1]) < 0 {
			grd = false
		}
		if f(a[i], a[i-1]) > 0 {
			min = false
		}
	}
	return grd || min
}
