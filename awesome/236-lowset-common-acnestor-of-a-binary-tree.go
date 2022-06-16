package awesome

import (
	"fmt"
	"runtime/debug"
)

//https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var  backRange func(*TreeNode) bool

	var parentRoot *TreeNode


	// 定义一个函数，使用后序遍历, 如果找到了任何一个p、或者q，则返回true
	backRange = func(node *TreeNode) bool {

		if node == nil {
			return false
		}
		fmt.Printf("node: %v \n", node.Val)
		foundLeft := backRange(node.Left)
		foundRight := backRange(node.Right)

		fmt.Printf("node: %v, left: %v, right: %v \n",
			node.Val, foundLeft, foundRight)
		// 说明在它的子节点
		if (node == p || node == q)  && (foundLeft && foundRight) {
			fmt.Printf("node: %v \n pick", node.Val)
			if parentRoot == nil {
				parentRoot = node
			}
		}
		// 说明在左右分支
		if foundLeft && foundRight {
			fmt.Printf("node: %v \n pick", node.Val)
			if parentRoot == nil {
				parentRoot = node
			}
		}
		if node == p || node == q {
			return true
		}

		return foundRight || foundLeft
	}

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("recover: %v \n", e)
			debug.PrintStack()
		}
	}()

	backRange(root)
	if parentRoot != nil {
		fmt.Printf("parentRoot: %v \n pick", parentRoot.Val)
	} else {
		fmt.Printf("parentRoot")
	}

	return parentRoot
}