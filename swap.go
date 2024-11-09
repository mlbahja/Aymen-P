package piscine

func Swap(a *int, b *int) {
	k := *a
	*a = *b
	*b = k
}
