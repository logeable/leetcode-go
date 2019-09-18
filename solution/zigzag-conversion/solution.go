package solution

import "bytes"

func convert(s string, numRows int) string {
	z := zigzag(s, numRows)
	var buf bytes.Buffer
	for _, r := range z {
		for _, c := range r {
			if c != 0 {
				buf.WriteRune(c)
			}
		}
	}
	return buf.String()
}

func zigzag(s string, numRows int) [][]rune {
	if numRows == 0 {
		return nil
	}
	if numRows == 1 {
		return [][]rune{[]rune(s)}
	}
	result := make([][]rune, numRows)
	sc := numRows - 2 + 1
	l := 2*numRows - 2
	col := len(s) / l
	rem := len(s) % l
	if rem != 0 {
		if rem <= numRows {
			rem = 1
		} else {
			rem = 1 + rem - numRows
		}
	}
	c := col*sc + rem

	for i := 0; i < numRows; i++ {
		result[i] = make([]rune, c)
		for j := 0; j < c; j++ {
			result[i][j] = 0
		}
	}

	for i, c := range s {
		m := i % l
		d := i / l
		var row, col int
		if m < numRows {
			row = m
			col = (numRows - 1) * d
		} else {
			col = (numRows-1)*d + (m - numRows + 1)
			row = 2*numRows - m - 2
		}
		result[row][col] = c
	}

	return result
}
