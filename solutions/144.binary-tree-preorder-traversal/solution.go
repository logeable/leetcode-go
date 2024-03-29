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
func preorderTraversal(root *TreeNode) []int {
	var arr []int
	pre(&arr, root)
	return arr
}

func pre(arr *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	pre(arr, root.Left)
	pre(arr, root.Right)
}
