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

	level := -1
	var ret [][]int
	var arr []int
	q := []*levelNode{{level: 0, node: root}}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.level > level {
			if level != -1 {
				ret = append(ret, arr)
			}
			level = node.level
			arr = nil
		}
		arr = append(arr, node.node.Val)
		if node.node.Left != nil {
			q = append(q, &levelNode{level: level + 1, node: node.node.Left})
		}
		if node.node.Right != nil {
			q = append(q, &levelNode{level: level + 1, node: node.node.Right})
		}
	}
	ret = append(ret, arr)

	return ret
}

type levelNode struct {
	level int
	node  *types.TreeNode
}
