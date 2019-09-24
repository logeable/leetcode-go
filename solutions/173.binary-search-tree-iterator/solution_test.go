package solution

import (
	"fmt"
	"testing"

	"github.com/logeable/leetcode-go/types"
	"github.com/stretchr/testify/assert"
)

func TestBSTIterator(t *testing.T) {
	table := []struct {
		input []interface{}
	}{
		{
			[]interface{}{7, 3, 15, nil, nil, 9, 20},
		},
	}
	for _, row := range table {
		root := types.MakeTreeNode(row.input)

		var inorderResult []int
		in(&inorderResult, root)

		iter := Constructor(root)
		var iterResult []int
		for iter.HasNext() {
			iterResult = append(iterResult, iter.Next())
		}

		assert.Equal(t, inorderResult, iterResult, fmt.Sprintf("input: %v\n%v", row.input, root))
	}
}

func in(arr *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	in(arr, root.Left)
	*arr = append(*arr, root.Val)
	in(arr, root.Right)
}
