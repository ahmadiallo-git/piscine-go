package z01

func IsPrintable(s string) bool {
	tabS := []rune(s)
	for i := 0; i < len(tabS); i++ {
		if tabS[i] < 32 && tabS[i] >= 0 {
			return false
		}
	}
	return true
}
