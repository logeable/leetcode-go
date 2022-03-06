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

	return sym(root.Left, root.Right)
}

func sym(a *types.TreeNode, b *types.TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a != nil && b != nil {
		return a.Val == b.Val && sym(a.Left, b.Right) && sym(a.Right, b.Left)
	} else {
		return false
	}
}
