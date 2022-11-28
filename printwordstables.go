package z01

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	for _, word := range a {
		myString := word
		for _, i := range myString {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')

	}
}
