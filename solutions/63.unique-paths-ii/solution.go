package solution

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	if obstacleGrid[0][0] == 1 {
		dp[0][0] = 0
	}
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			dp[0][i] = 0
			continue
		}
		dp[0][i] = dp[0][i-1]
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
			continue
		}
		dp[i][0] = dp[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
