package piscine

func Compact(ptr *[]string) int {
	slc := *ptr
	count := 0
	for i := 0; i < len(slc); i++ {
		if slc[i] != "" {
			slc[count] = slc[i]
			count++
		}
	}
	*ptr = slc[:count]
	return count
}
