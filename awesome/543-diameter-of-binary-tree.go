package awesome

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {


	maxPathLength := math.MinInt32

	var orderTree func(node *TreeNode) int
	orderTree = func(node *TreeNode) int {

		if node == nil {
			return 0
		}

		leftPath := orderTree(node.Left)
		rightPath := orderTree(node.Right)

		maxPath := max(leftPath,rightPath)

		maxPathLength = max(
			leftPath+rightPath,
			maxPathLength)

		return maxPath + 1
	}
	orderTree(root)
	return maxPathLength

}