package solution

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	index := make(map[byte]bool)
	ans := 1
	end := -1
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(index, s[i-1])
		}
		for end+1 < len(s) && !index[s[end+1]] {
			end++
			index[s[end]] = true
		}
		if v := end - i + 1; v > ans {
			ans = v
		}
	}
	return ans
}
