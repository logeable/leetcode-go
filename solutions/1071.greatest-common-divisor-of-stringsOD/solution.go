package solution

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return b
}
