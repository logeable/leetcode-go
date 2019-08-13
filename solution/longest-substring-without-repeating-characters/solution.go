package solution

func lengthOfLongestSubstringBruteForce(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			x := s[i:j]
			if uniq(x) && len(x) > res {
				res = len(x)
			}
		}
	}
	return res
}

func uniq(s string) bool {
	cache := make(map[rune]struct{})
	for _, x := range s {
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	index := make(map[rune]int)
	var res int
	var begin int
	for i, x := range s {
		if last, ok := index[x]; ok {
			begin = last + 1
		}
		index[x] = i
		if i-begin+1 > res {
			res = i - begin + 1
		}
	}
	return res
}
