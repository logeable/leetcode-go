package solution

func toLowerCase(str string) string {
	var res string
	for _, x := range str {
		res += string(x | 0b00100000)
	}
	return res
}
