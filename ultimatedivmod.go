package z01

func UltimateDivMod(a *int, b *int) {
	tempA := *a / *b
	tempB := *a % *b
	*a = tempA
	*b = tempB
}
