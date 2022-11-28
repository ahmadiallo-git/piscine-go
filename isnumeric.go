package z01

func IsNumeric(s string) bool {
	tabS := []rune(s)

	for i := 0; i < len(tabS); i++ {
		if tabS[i] < '0' || tabS[i] > '9' {
			return false
		}
	}
	return true
}
