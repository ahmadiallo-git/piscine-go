package z01

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				value := table[i]
				table[i] = table[j]
				table[j] = value
			}
		}
	}
}
