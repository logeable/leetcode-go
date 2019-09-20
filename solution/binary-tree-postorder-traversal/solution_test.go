package solution

import (
	"fmt"
	"testing"

	"github.com/logeable/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func Test_postorderTraversal(t *testing.T) {
	table := []struct {
		input    []interface{}
		expected []int
	}{
		{
			[]interface{}{},
			nil,
		},
		{
			[]interface{}{nil},
			nil,
		},
		{
			[]interface{}{1, nil, 2, 3},
			[]int{3, 2, 1},
		},
		{
			[]interface{}{1, nil, 2, 3, nil, nil, 4, 5, 6},
			[]int{5, 6, 4, 3, 2, 1},
		},
	}

	for _, row := range table {
		root := types.MakeTreeNode(row.input)
		actual := postorderTraversal(root)
		assert.Equal(t, row.expected, actual, fmt.Sprintf("input: %v, actual: %v", row.input, actual))
	}
}
