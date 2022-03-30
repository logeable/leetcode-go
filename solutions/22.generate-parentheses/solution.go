package solution

func generateParenthesis(n int) []string {
	dp := make([][]string, n)
	dp[0] = []string{"()"}
	for i := 1; i < n; i++ {
		cache := make(map[string]struct{})
		for _, x := range dp[i-1] {
			for _, v := range addParenthes(x) {
				cache[v] = struct{}{}
			}
		}
		for k := range cache {
			dp[i] = append(dp[i], k)
		}
	}
	return dp[n-1]
}

func addParenthes(p string) []string {
	var ret []string

	for i := 0; i < len(p); i++ {
		ret = append(ret, p[0:i]+"()"+p[i:])
	}
	ret = append(ret, p+"()")
	return ret
}
