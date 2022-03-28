package solution

func addBinary(a string, b string) string {
	var ret []byte
	i := len(a) - 1
	j := len(b) - 1

	c := byte(0)
	for i >= 0 || j >= 0 {
		t := c
		c = 0
		if i >= 0 {
			t += a[i] - '0'
			i--
		}
		if j >= 0 {
			t += b[j] - '0'
			j--
		}
		if t == 2 {
			c = 1
			t = 0
		} else if t == 3 {
			c = 1
			t = 1
		}
		ret = append([]byte{t + '0'}, ret...)
	}
	if c > 0 {
		ret = append([]byte{c + '0'}, ret...)
	}
	return string(ret)
}
