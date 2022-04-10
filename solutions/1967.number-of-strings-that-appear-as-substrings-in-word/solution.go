package solution

import "strings"

func numOfStrings(patterns []string, word string) int {
	ans := 0
	for _, p := range patterns {
		if strings.Index(word, p) != -1 {
			ans++
		}
	}
	return ans
}
