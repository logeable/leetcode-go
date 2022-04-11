package solution

func countNumbersWithUniqueDigits(n int) int {
	dp := make([]int, n+2)
	dp[0] = 1
	dp[1] = 10
	if n <= 1 {
		return dp[n]
	}
	for i := 2; i <= n; i++ {
		t := 9
		for j := 0; j < i-1; j++ {
			t *= 9 - j
		}
		t += dp[i-1]
		dp[i] = t
	}
	return dp[n]
}
