package piscine

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	str := []rune(Itoa(n))
	Sort(str)
	PrintStr(string(str))
}

func Itoa(n int) string {
	var str string
	if n == 0 {
		str = "0"
	}
	for n > 0 {
		str += string(rune(n%10) + '0')
		n /= 10
	}
	return str
}

func Sort(n []rune) string {
	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	return string(n)
}
