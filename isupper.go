package z01

func IsUpper(s string) bool {
	tabS := []rune(s)
	for i := 0; i < len(tabS); i++ {
		if tabS[i] < 'A' || tabS[i] > 'Z' {
			return false
		}
	}
	return true
}
