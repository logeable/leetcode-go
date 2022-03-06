package solution

import "github.com/logeable/leetcode-go/types"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *types.TreeNode) bool {
	if root == nil {
		return true
	}
	q := []*types.TreeNode{root}

	for len(q) != 0 {
		var layer []*types.TreeNode
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]

			layer = append(layer, node)
			if node != nil {
				q = append(q, node.Left)
				q = append(q, node.Right)
			}
		}
		l, r := 0, length-1
		for l < r {
			if layer[l] == nil && layer[r] == nil {
				l++
				r--
			} else if layer[l] != nil && layer[r] != nil && layer[l].Val == layer[r].Val {
				l++
				r--
			} else {
				return false
			}
		}

	}
	return true
}
