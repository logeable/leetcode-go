package solution

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	start := 0
	length := 1
	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s)-l+1; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				if j-i == 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = true && dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > length {
				start = i
				length = l
			}
		}
	}

	return s[start : start+length]
}
