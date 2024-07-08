package solution

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	mini := make([]int, len(triangle))
	copy(mini, triangle[len(triangle)-1])

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			mini[j] = triangle[i][j] + min(mini[j], mini[j+1])
		}
	}
	return mini[0]
}

func traverse(triangle [][]int, i, j int) int {
	if i == len(triangle)-1 {
		return triangle[i][j]
	}
	return triangle[i][j] + min(traverse(triangle, i+1, j), traverse(triangle, i+1, j+1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
