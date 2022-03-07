package solution

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([]int, len(prices))
	dp[0] = 0

	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		dp[i] = dp[i-1]
		v := prices[i] - min
		if v > dp[i-1] {
			dp[i] = v
		}
	}

	return dp[len(prices)-1]
}
