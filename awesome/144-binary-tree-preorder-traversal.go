package awesome

import "fmt"

// https://leetcode.cn/problems/binary-tree-preorder-traversal/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
func preorderTraversal(root *TreeNode) []int {

	// 答案
	var ans []int
	var rePreOrderTraversal func(root *TreeNode)
	rePreOrderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		fmt.Printf("node: %v", node.Val)
		rePreOrderTraversal(node.Left)
		rePreOrderTraversal(node.Right)
	}
	rePreOrderTraversal(root)
	return ans
}


func preorderTraversal2(root *TreeNode) []int {
	// 答案
	var ans []int
	if root == nil {
		return ans
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for ;len(stack) > 0; {

		top := stack[len(stack)-1] //栈顶
		stack = stack[:len(stack)-1]
		ans = append(ans, top.Val)
		fmt.Printf("cur: %v, stack len: %v \n", top.Val, len(stack))

		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}

	}
	return ans
}