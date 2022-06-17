package awesome

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。
//
//提醒一下，二叉搜索树满足下列约束条件：
//
//节点的左子树仅包含键 小于 节点键的节点。
//节点的右子树仅包含键 大于 节点键的节点。
//左右子树也必须是二叉搜索树。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/convert-bst-to-greater-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func convertBST(root *TreeNode) *TreeNode {
	//  后序遍历
	var backOrderTree func(node *TreeNode)

	// 累加值
	sum := 0
	backOrderTree = func(node *TreeNode) {

		if node == nil {
			return
		}

		backOrderTree(node.Right)
		fmt.Printf("node: %v, sum: %v \n", node.Val, sum)
		temp := node.Val
		node.Val = node.Val + sum
		sum = sum + temp
		backOrderTree(node.Left)
	}

	backOrderTree(root)
	return root
}