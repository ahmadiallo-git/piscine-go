package z01

func IsAlpha(s string) bool {
	tabS := []rune(s)
	for i := 0; i < len(tabS); i++ {
		if tabS[i] < '0' || tabS[i] > '9' && tabS[i] < 'A' || tabS[i] > 'Z' && tabS[i] < 'a' || tabS[i] > 'z' {
			return false
		}
	}
	return true
}
