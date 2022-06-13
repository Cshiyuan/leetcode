package awesome

import "fmt"

//给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
//展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
//展开后的单链表应该与二叉树 先序遍历 顺序相同。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/flatten-binary-tree-to-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	newRoot := &TreeNode{
		Val: root.Val,
	}
	tree := preRange(root.Left, newRoot)
	preRange(root.Right, tree)
	(*root) = *newRoot
	fmt.Print(root)
}

func preRange(root *TreeNode, newTree *TreeNode) *TreeNode {
	if root == nil {
		return newTree
	}
	newTree.Right = &TreeNode{
		Val: root.Val,
	}
	newTree = newTree.Right
	// 先左边
	tree := preRange(root.Left, newTree)
	tree = preRange(root.Right, tree)
	return tree
}
