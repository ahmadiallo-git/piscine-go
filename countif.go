package z01

func CountIf(f func(string) bool, tab []string) int {
	nbrel := 0
	for _, s := range tab {
		if f(s) == true {
			nbrel++
		}
	}
	return nbrel
}
