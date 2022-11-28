package z01

func MakeRange(min, max int) []int {
	var answer []int
	if min >= max {
		answer = nil
	} else {
		answer = make([]int, max-min)

		for i := 0; i < (max - min); i++ {
			answer[i] = i + min
		}
		return answer
	}
	return answer
}
