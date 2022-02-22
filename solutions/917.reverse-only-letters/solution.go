package solution

func reverseOnlyLetters(s string) string {
	runes := []rune(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isLetter(runes[l]) {
			l++
		}
		for l < r && !isLetter(runes[r]) {
			r--
		}
		if l >= r {
			break
		}
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	return string(runes)
}

func isLetter(r rune) bool {
	return (r|32) >= 'a' && (r|32) <= 'z'
}
