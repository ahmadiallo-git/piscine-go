package z01

func SplitWhiteSpaces(s string) []string {
	var mySpace []string
	var result string

	for i, n := range s {
		if n > 32 {
			result += string(n)
		} else if result != "" {
			mySpace = append(mySpace, result)
			result = ""

		}
		if i == len(s)-1 {
			mySpace = append(mySpace, result)
		}

	}
	return mySpace
}
