package z01

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}
