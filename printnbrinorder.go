package z01

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	var s [10]int
	if n == 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		s[n%10]++
		n = n / 10

	}
	for i := 0; i < 10; i++ {
		for s[i] > 0 {
			z01.PrintRune(rune(i) + '0')
			s[i]--
		}
	}
}
