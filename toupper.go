package z01

func ToUpper(s string) string {
	tab := []rune(s)

	for i := 0; i <= len(tab)-1; i++ {
		if tab[i] >= 'a' && tab[i] <= 'z' {
			tab[i] -= 32
		}
	}
	return string(tab)
}
