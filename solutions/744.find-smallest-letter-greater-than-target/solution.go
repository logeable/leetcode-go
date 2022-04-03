package solution

func nextGreatestLetter(letters []byte, target byte) byte {
	if target >= letters[len(letters)-1] {
		return letters[0]
	}

	l, r := 0, len(letters)-1
	for l < r {
		m := l + (r-l)>>1
		if letters[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return letters[l]
}
