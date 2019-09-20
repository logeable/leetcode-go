package solution

import (
	"fmt"
	"testing"

	"github.com/logeable/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func Test_preorderTraversal(t *testing.T) {
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
			[]int{1, 2, 3},
		},
		{
			[]interface{}{1, nil, 2, 3, nil, nil, 4, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, row := range table {
		root := types.MakeTreeNode(row.input)
		actual := preorderTraversal(root)
		assert.Equal(t, row.expected, actual, fmt.Sprintf("input: %v, actual: %v", row.input, actual))
	}
}
