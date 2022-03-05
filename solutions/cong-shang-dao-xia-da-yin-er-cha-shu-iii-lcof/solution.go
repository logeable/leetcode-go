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
func levelOrder(root *types.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	q := []*types.TreeNode{root}

	level := 0
	for len(q) != 0 {
		length := len(q)
		var layer []int
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]
			if level&1 == 0 {
				layer = append(layer, node.Val)
			} else {
				layer = append([]int{node.Val}, layer...)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ret = append(ret, layer)
		level++
	}

	return ret

}
