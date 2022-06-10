package awesome

import "math"

//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
//有效 二叉搜索树定义如下：
//
//节点的左子树只包含 小于 当前节点的数。
//节点的右子树只包含 大于 当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。

func isValidBST(root *TreeNode) bool {
	b, _, _ := reIsValidBST(root)
	return b
}

// 返回最大最小
func reIsValidBST(root *TreeNode) (bool, int, int) {
	if root == nil {
		return true, math.MaxInt64, math.MinInt64
	}
	if root.Left != nil && root.Left.Val >= root.Val {
		return false, 0, 0
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false, 0, 0
	}
	// 叶子节点
	if root.Right == nil && root.Left == nil {
		return true, root.Val, root.Val
	}
	//var rightIsBST
	rightIsBST, rMin, rMax := reIsValidBST(root.Right)
	leftIsBST, lMin, lMax := reIsValidBST(root.Left)

	// 两边都是平衡二叉树
	if rightIsBST && leftIsBST {
		//如果右边的最小值,小于等于根
		if rMin <= root.Val  {
			return false, 0, 0
		}
		// 左边的最大值，大于跟
		if lMax >= root.Val  {
			return false, 0, 0
		}
		return true, lMin, rMax
	}

	return false, 0, 0
}
