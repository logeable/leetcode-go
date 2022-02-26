package solution

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := 0
	max := nums[0]
	for _, n := range nums {
		sum += n
		if n > max {
			max = n
		}
	}
	if sum&1 == 1 {
		return false
	}
	target := sum / 2
	if max > target {
		return false
	}

	dp := make([][]bool, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	for i := 1; i <= target; i++ {
		dp[0][i] = nums[0] == target
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)-1][target]
}
