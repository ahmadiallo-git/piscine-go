package z01

func LastRune(s string) rune {
	res := []rune(s)
	arr := 0
	for i := range s {
		arr = i + 1
	}
	return res[arr-1]
}
