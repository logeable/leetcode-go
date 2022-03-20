package solution

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}

	ret := make([][]int, numRows)
	ret[0] = []int{1}
	ret[1] = []int{1, 1}

	for i := 2; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		ret[i][0], ret[i][i] = 1, 1
		for j := 1; j < i; j++ {
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
	}

	return ret[:numRows]
}
