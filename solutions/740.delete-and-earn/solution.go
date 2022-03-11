package solution

func deleteAndEarn(nums []int) int {
	sum := make([]int, 10001)
	for _, n := range nums {
		sum[n] += n
	}

	dp := make([]int, len(sum))
	dp[0] = sum[0]
	dp[1] = max(sum[0], sum[1])
	for i := 2; i < len(sum); i++ {
		dp[i] = max(dp[i-2]+sum[i], dp[i-1])
	}
	return dp[10000]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
