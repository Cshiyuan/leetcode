package awesome

import (
	"fmt"
	"math"
)

//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//
//路径和 是路径中各节点值的总和。
//
//给你一个二叉树的根节点 root ，返回其 最大路径和 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/binary-tree-maximum-path-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {

		if node == nil {
			return 0
		}
		// 递归计算左右节点最大的贡献值
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		// 假如当前节点用作路径的根节点
		curGain := leftGain + rightGain + node.Val

		maxSum = max(maxSum, curGain)
		maxSum = max(maxSum, node.Val)
		fmt.Printf("nodeVal: %v, curGain: %v, maxSum: %v \n", node.Val, curGain, maxSum)
		return node.Val + max(leftGain, rightGain)
	}

	rootVal := maxGain(root)
	maxSum = max(maxSum, rootVal)

	return maxSum
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}