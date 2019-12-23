package solution

import (
	"fmt"
	"testing"

	"github.com/logeable/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {
	table := []struct {
		input    []interface{}
		expected [][]int
	}{
		{
			[]interface{}{},
			nil,
		},
		{
			[]interface{}{3, 9, 20, nil, nil, 15, 7},
			[][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			[]interface{}{1, 2, 3, 4, 5, 6, 7, 8},
			[][]int{
				{1},
				{2, 3},
				{4, 5, 6, 7},
				{8},
			},
		},
		{
			[]interface{}{1, 2, 3, nil, 4, 5, nil, 6, nil, nil, 7, nil, nil, nil, 8},
			[][]int{
				{1},
				{2, 3},
				{4, 5},
				{6, 7},
				{8},
			},
		},
	}

	for _, row := range table {
		input := types.MakeTreeNode(row.input)
		actual := levelOrder(input)
		assert.Equal(t, row.expected, actual, fmt.Sprintf("input:\n%v\nactual: %d", input, actual))
	}
}
