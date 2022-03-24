package solution

func imageSmoother(img [][]int) [][]int {
	rows, cols := len(img), len(img[0])

	ret := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ret[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			ret[i][j] = calc(img, i, j)
		}
	}
	return ret
}

func calc(img [][]int, r, c int) int {
	rows, cols := len(img), len(img[0])
	total := 0
	count := 0

	x, y := r-1, c-1
	if valid(x, y, rows, cols) {
		total += img[x][y]
		count++
	}
	x, y = r-1, c
	if valid(x, y, rows, cols) {
		total += img[x][y]
		count++
	}
	x, y = r-1, c+1
	if valid(x, y, rows, cols) {
		total += img[x][y]
		count++
	}
	x, y = r, c-1
	if valid(x, y, rows, cols) {
		total += img[x][y]
		count++
	}
	x, y = r, c
	if valid(x, y, rows, cols) {
		total += img[x][y]
		count++
	}
	x, y = r, c+1
	if valid(x, y, rows, cols) {
		total += img[x][y]
		count++
	}
	x, y = r+1, c-1
	if valid(x, y, rows, cols) {
		total += img[x][y]
		count++
	}
	x, y = r+1, c
	if valid(x, y, rows, cols) {
		total += img[x][y]
		count++
	}
	x, y = r+1, c+1
	if valid(x, y, rows, cols) {
		total += img[x][y]
		count++
	}
	return total / count
}

func valid(r, c int, rows, cols int) bool {
	return r >= 0 && c >= 0 && r < rows && c < cols
}
