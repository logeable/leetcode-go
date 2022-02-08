package solution

func reverseLeftWords(s string, n int) string {
	rs := []rune(s)
	reverse(rs[:n])
	reverse(rs[n:])
	reverse(rs)
	return string(rs)
}

func reverse(rs []rune) {
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
}
