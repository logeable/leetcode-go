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
func inorderTraversal(root *types.TreeNode) []int {
	var arr []int
	in(&arr, root)
	return arr
}

func in(arr *[]int, root *types.TreeNode) {
	if root == nil {
		return
	}
	in(arr, root.Left)
	*arr = append(*arr, root.Val)
	in(arr, root.Right)
}
