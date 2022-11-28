package z01

import "github.com/01-edu/z01"

func safe(a int, b int, pos [8]int) bool {
	for i := 0; i < a; i++ {
		t := pos[i]
		if t == b || t == b-(a-i) || t == b+(a-i) {
			return false
		}
	}
	return true
}

func resoudre(nb int, pos [8]int) {
	if nb == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune(pos[i] + 49))
		}
		z01.PrintRune(10)
	} else {
		for i := 0; i < 8; i++ {
			if safe(nb, i, pos) {
				pos[nb] = i
				resoudre(nb+1, pos)
			}
		}
	}
}

func EightQueens() {
	pos := [8]int{0, 0, 0, 0, 0, 0, 0}

	resoudre(0, pos)
}
