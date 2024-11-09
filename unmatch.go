package piscine

func Unmatch(a []int) int {
	for l := 0; l < len(a); l++ {
		count := 0
		k := a[l]
		for h := 0; h < len(a); h++ {
			if k == a[h] {
				count++
			}
		}
		if count%2 == 1 {
			return k
		}
	}
	return -1
}
