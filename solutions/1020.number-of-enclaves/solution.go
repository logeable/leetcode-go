package solution

func numEnclaves(grid [][]int) int {
	var total int
	var roots [][2]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && (i == 0 || i == len(grid)-1 || j == 0 || j == len(grid[0])-1) {
				roots = append(roots, [2]int{i, j})
			}
			if grid[i][j] == 1 {
				total++
			}
		}
	}

	visit := make([][]bool, len(grid))
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, len(grid[0]))
		for j := 0; j < len(visit[0]); j++ {
			visit[i][j] = false
		}
	}

	ret := total
	for _, p := range roots {
		r, c := p[0], p[1]
		if visit[r][c] {
			continue
		}
		n := linkNodes(grid, visit, r, c)
		ret -= n
	}
	return ret
}

func linkNodes(grid [][]int, visit [][]bool, r, c int) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return 0
	}
	if visit[r][c] {
		return 0
	}
	visit[r][c] = true

	if grid[r][c] == 1 {
		return 1 + linkNodes(grid, visit, r-1, c) + linkNodes(grid, visit, r+1, c) + linkNodes(grid, visit, r, c-1) + linkNodes(grid, visit, r, c+1)
	}
	return 0
}
