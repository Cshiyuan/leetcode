package awesome

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
//
//除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，
//聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
//如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
//
//给定二叉树的 root 。
//返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。

func rob_iii(root *TreeNode) int {
	// 中序遍历

	cacheResult := make(map[*TreeNode]int, 0)
	// cacheUnPickResult := make(map[string]int, 0)

	var treeMax func(node *TreeNode) int
	treeMax = func(node *TreeNode) int {


		val, ok := cacheResult[node]
		if ok {
			return val
		}

		if node == nil {
			return 0
		}
		if node.Right == nil && node.Left == nil {
			cacheResult[node] = node.Val
			return node.Val
		}

		pickLeftMax := 0
		if node.Left != nil {
			pickLeftMax = treeMax(node.Left.Left) + treeMax(node.Left.Right)
		}
		pickRightMax := 0
		if node.Right != nil {
			pickRightMax = treeMax(node.Right.Left) + treeMax(node.Right.Right)
		}

		pickMax := node.Val + pickLeftMax + pickRightMax
		unPickMax := treeMax(node.Right) + treeMax(node.Left)

		val = max(pickMax, unPickMax)
		cacheResult[node] = val
		return val
	}
	return treeMax(root)
}

// 其他解法
func rob_iii2(root *TreeNode) int {
	// 中序遍历
	var treeMax func(node *TreeNode) (int, int)
	treeMax = func(node *TreeNode) (int, int) {

		if node == nil {
			return 0, 0
		}
		if node.Right == nil && node.Left == nil {
			return node.Val, 0
		}
		rPick, rUnPick := treeMax(node.Right)
		lPick, lUnPick := treeMax(node.Left)

		pickMax := node.Val + rUnPick + lUnPick
		unPickMax := max(rPick, rUnPick) + max(lPick, lUnPick)

		return pickMax, unPickMax
	}

	pickMax, unPickMax := treeMax(root)

	return max(pickMax, unPickMax)
}