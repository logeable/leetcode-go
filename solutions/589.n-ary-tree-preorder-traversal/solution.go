package solution

import "github.com/logeable/leetcode-go/types"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *types.Node) []int {
	if root == nil {
		return nil
	}
	ret := []int{root.Val}
	for i := 0; i < len(root.Children); i++ {
		ret = append(ret, preorder(root.Children[i])...)
	}

	return ret
}
