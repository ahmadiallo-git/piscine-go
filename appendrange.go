package z01

func AppendRange(min, max int) []int {
	var answer []int

	if min >= max {
		answer = (nil)
	} else {
		for i := min; i < max; i++ {
			answer = append(answer, i)
		}
		return answer
	}
	return answer
}
