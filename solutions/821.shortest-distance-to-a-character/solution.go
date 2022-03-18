package solution

import "math"

func shortestToChar(s string, c byte) []int {
	ret := make([]int, len(s))
	prev := -math.MaxInt / 2
	for i := 0; i < len(ret); i++ {
		if s[i] == c {
			prev = i
		}
		ret[i] = i - prev
	}
	prev = math.MaxInt

	for i := len(ret) - 1; i >= 0; i-- {
		if s[i] == c {
			prev = i
		}
		ret[i] = min(ret[i], prev-i)
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
