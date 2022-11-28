package z01

func Swap(a *int, b *int) {
	temp1 := *a
	*a = *b
	*b = temp1
}
