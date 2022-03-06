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
func isSubStructure(A *types.TreeNode, B *types.TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return hasSubStructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func hasSubStructure(A *types.TreeNode, B *types.TreeNode) bool {
	if A == nil && B == nil {
		return true
	} else if A != nil && B == nil {
		return true
	} else if A == nil && B != nil {
		return false
	} else if A != nil && B != nil {
		return A.Val == B.Val && hasSubStructure(A.Left, B.Left) && hasSubStructure(A.Right, B.Right)
	}
	return false
}
