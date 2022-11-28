package z01

func Index(s string, toFind string) int {
	tabS := []rune(s)
	tabFind := []rune(toFind)
	lens := len(tabS)
	lenF := len(tabFind)

	for i := 0; i <= lens-lenF; i++ {
		if toFind == s[i:i+lenF] {
			return (i)
		}
	}
	return -1
}
