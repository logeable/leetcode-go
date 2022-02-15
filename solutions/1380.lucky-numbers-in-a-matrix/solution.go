package solution

func luckyNumbers(matrix [][]int) []int {
	var ret []int
	column := make([]int, len(matrix))
	for _, row := range matrix {
		mins := findMin(row)
		for _, i := range mins {
			fillColumn(matrix, i, column)
			if isMax(column, row[i]) {
				ret = append(ret, row[i])
			}
		}

	}
	return ret
}

func findMin(nums []int) []int {
	var ret []int
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	for i, v := range nums {
		if v == min {
			ret = append(ret, i)
		}
	}
	return ret
}

func isMax(nums []int, v int) bool {
	for _, n := range nums {
		if n > v {
			return false
		}
	}
	return true
}

func fillColumn(matrix [][]int, c int, nums []int) {
	for i, row := range matrix {
		nums[i] = row[c]
	}
}
