package awesome

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// https://leetcode.cn/problems/symmetric-tree/
func isSymmetric(root *TreeNode) bool {
	return rIsSymmetric(root.Left, root.Right)
}

func rIsSymmetric(leftRoot *TreeNode, rightRoot *TreeNode) bool {

	if leftRoot == nil && leftRoot == rightRoot {
		return true
	}
	if leftRoot == nil && rightRoot != nil {
		return false
	}
	if leftRoot != nil && rightRoot == nil {
		return false
	}

	if leftRoot.Val != rightRoot.Val {
		return false
	}
	return rIsSymmetric(leftRoot.Left, rightRoot.Right) &&
		rIsSymmetric(leftRoot.Right, rightRoot.Left)
}
