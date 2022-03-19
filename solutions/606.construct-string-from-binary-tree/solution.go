package solution

import (
	"fmt"

	"github.com/logeable/leetcode-go/types"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(root *types.TreeNode) string {
	if root == nil {
		return ""
	}

	if root.Left != nil && root.Right != nil {
		return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
	} else if root.Left != nil {
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	} else if root.Right != nil {
		return fmt.Sprintf("%d()(%s)", root.Val, tree2str(root.Right))
	} else {
		return fmt.Sprintf("%d", root.Val)
	}
}
