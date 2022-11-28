package z01

func IsLower(s string) bool {
	tabS := []rune(s)

	for i := 0; i < len(tabS); i++ {
		if tabS[i] < 'a' || tabS[i] > 'z' {
			return false
		}
	}

	return true
}
