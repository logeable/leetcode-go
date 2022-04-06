package solution

func permutation(s string) []string {
	var ans []string
	cache := make(map[string]struct{})
	ret := perm(s, 0)
	for _, v := range ret {
		cache[v] = struct{}{}
	}
	for k := range cache {
		ans = append(ans, k)
	}
	return ans
}

func perm(s string, begin int) []string {
	var ret []string
	if begin == len(s) {
		ret = append(ret, s)
	}
	bs := []byte(s)
	for i := begin; i < len(s); i++ {
		bs[begin], bs[i] = bs[i], bs[begin]
		ret = append(ret, perm(string(bs), begin+1)...)
		bs[begin], bs[i] = bs[i], bs[begin]
	}

	return ret
}
