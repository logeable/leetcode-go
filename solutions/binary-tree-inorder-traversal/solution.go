package solution

import "github.com/logeable/leetcode-go/types"

type TreeNode = types.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var arr []int
	in(&arr, root)
	return arr
}

func in(arr *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	in(arr, root.Left)
	*arr = append(*arr, root.Val)
	in(arr, root.Right)
}
