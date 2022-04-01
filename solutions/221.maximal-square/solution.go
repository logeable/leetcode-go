package solution

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	max := 0
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(dp[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				max = 1
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j-1], dp[i][j-1]), dp[i-1][j]) + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max * max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
