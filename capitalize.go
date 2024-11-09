package piscine

func Capitalize(s string) string {
	str := ""
	capitalizeNext := true

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if capitalizeNext {
				if char >= 'a' && char <= 'z' {
					str += string(char - 32)
				} else {
					str += string(char)
				}
				capitalizeNext = false
			} else {
				if char >= 'A' && char <= 'Z' {
					str += string(char + 32)
				} else {
					str += string(char)
				}
			}
		} else {
			str = str + string(char)
			capitalizeNext = true
		}
	}

	return str
}
