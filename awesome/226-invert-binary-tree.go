package awesome

//给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归反转
func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	var oLeft *TreeNode
	oLeft = root.Left

	root.Left = invertTree(root.Right)

	root.Right = invertTree(oLeft)

	return root
}
