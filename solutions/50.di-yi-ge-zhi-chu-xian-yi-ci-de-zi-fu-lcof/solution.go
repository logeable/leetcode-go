package solution

func firstUniqChar(s string) byte {
	var cache [26]int
	for i := range s {
		cache[s[i]-'a']++
	}
	for i := range s {
		if cache[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
