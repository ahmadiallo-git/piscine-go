package z01

func StrLen(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if byte(s[i]) < 129 {
			count++
		}
	}
	word := len(s)
	if count != word {
		count++
	}
	return count
}
