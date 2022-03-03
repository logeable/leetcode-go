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
func levelOrder(root *types.TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	q := []*types.TreeNode{root}
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		ret = append(ret, cur.Val)
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	return ret
}
