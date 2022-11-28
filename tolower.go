package z01

func ToLower(s string) string {
	tab := []rune(s)

	for i := 0; i <= len(tab)-1; i++ {
		if tab[i] >= 'A' && tab[i] <= 'Z' {
			tab[i] += 32
		}
	}
	return string(tab)
}
