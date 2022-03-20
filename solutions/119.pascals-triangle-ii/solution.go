package solution

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	p := []int{1}
	var q []int
	for i := 1; i < rowIndex+1; i++ {
		q = make([]int, i+1)
		q[0], q[i] = 1, 1
		for j := 1; j < i; j++ {
			q[j] = p[j-1] + p[j]
		}
		p = q
	}

	return q
}
