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
func mirrorTree(root *types.TreeNode) *types.TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Right
	root.Right = mirrorTree(root.Left)
	root.Left = mirrorTree(tmp)
	return root
}
