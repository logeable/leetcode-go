package solution

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		ret[i][0], ret[i][i] = 1, 1
		for j := 1; j < i; j++ {
			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
	}
	return ret
}
