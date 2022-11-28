package z01

func Atoi(s string) int {
	Enter := []rune(s)
	output := 0
	t := len(s)
	var signe int = 1

	if s != "" {
		if byte(Enter[0]) == 43 { // Desc +
			Enter[0] = '0'
		} else if Enter[0] == '-' { // Desc - the man assci
			signe *= -1
			Enter[0] = '0'
		}
	}

	for i := 0; i < t; i++ {
		switch {
		case byte(Enter[i]) < 48:
			return 0
			break
		case byte(Enter[i]) > 57:
			return 0
			break
		case byte(Enter[i]) == 32: // DESC SPACE
			return 0
		default:
			output = output*10 + int(Enter[i]) - '0'

		}
	}
	return output * signe
}
