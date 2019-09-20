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
func postorderTraversal(root *types.TreeNode) []int {
	var arr []int
	post(&arr, root)
	return arr
}

func post(arr *[]int, root *types.TreeNode) {
	if root == nil {
		return
	}
	post(arr, root.Left)
	post(arr, root.Right)
	*arr = append(*arr, root.Val)
}
