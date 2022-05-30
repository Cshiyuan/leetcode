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
//给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
//
// 
//
//示例 :
//给定二叉树
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/diameter-of-binary-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func diameterOfBinaryTree(root *TreeNode) int {

	val := math.MinInt64
	leftLength := rDiameterOfBinaryTree(root.Left, &val)
	rightLength := rDiameterOfBinaryTree(root.Right, &val)
	if val > rightLength + leftLength {
		return val
	}
	return rightLength + leftLength
}

func rDiameterOfBinaryTree(root *TreeNode, val *int) int {
	if root == nil {
		return 0
	}
	leftLength := rDiameterOfBinaryTree(root.Left, val)
	rightLength := rDiameterOfBinaryTree(root.Right, val)

	if leftLength+rightLength > *val {
		*val = (leftLength + rightLength)
	}
	maxLength := leftLength
	if rightLength > maxLength {
		maxLength = rightLength
	}
	return maxLength + 1
}
