package solution

import "math"

func convert(s string, numRows int) string {
	if numRows == 1 || numRows == len(s) {
		return s
	}
	table := make([][]byte, numRows)
	count := numRows*2 - 2
	times := int(math.Ceil(float64(len(s)) / float64(count)))
	cols := times * (numRows - 1)
	for i := range table {
		table[i] = make([]byte, cols)
	}

	r, c := 0, 0
	for i, b := range s {
		table[r][c] = byte(b)
		if i%count < numRows-1 {
			r++
		} else {
			r--
			c++
		}
	}

	ret := make([]byte, 0, len(s))
	for _, row := range table {
		for _, b := range row {
			if b > 0 {
				ret = append(ret, b)
			}
		}
	}
	return string(ret)
}
