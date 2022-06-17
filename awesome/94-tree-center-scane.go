package awesome

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
func inorderTraversal(root *TreeNode) []int {
	return rOrderTraversal(root)
}

func rOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	leftSlice := rOrderTraversal(root.Left)
	leftSlice = append(leftSlice, root.Val)
	rightSlice := rOrderTraversal(root.Right)
	return append(leftSlice, rightSlice...)
}
