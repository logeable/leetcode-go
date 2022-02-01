package solution

func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	col := len(matrix[0]) - 1

	for row < len(matrix) && col >= 0 {
		v := matrix[row][col]
		switch {
		case v == target:
			return true
		case v > target:
			col--
		case v < target:
			row++
		}
	}
	return false
}
