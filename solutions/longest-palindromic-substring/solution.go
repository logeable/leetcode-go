package solution

func longestPalindromeBruteForce(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			t := s[i:j]
			if isPalindrome(t) && len(t) > len(res) {
				res = t
			}
		}
	}
	return res
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func longestPalindrome(s string) string {
	panic("not implemented")
}
