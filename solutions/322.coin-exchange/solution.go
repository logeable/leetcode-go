package solution

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		count := math.MaxInt
		for _, coin := range coins {
			if coin <= i {
				if dp[i-coin] == -1 {
					continue
				}
				count = min(count, dp[i-coin])
				dp[i] = count + 1
			}
		}
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
