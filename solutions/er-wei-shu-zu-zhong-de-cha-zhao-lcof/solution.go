package solution

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	if len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])
	r, c := 0, cols-1
	for r < rows && c >= 0 {
		v := matrix[r][c]
		if v > target {
			c -= 1
		} else if v < target {
			r += 1
		} else {
			return true
		}
	}
	return false
}
