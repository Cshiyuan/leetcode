package awesome

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
//每条从根节点到叶节点的路径都代表一个数字：
//
//例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
//计算从根节点到叶节点生成的 所有数字之和 。
//
//叶节点 是指没有子节点的节点。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/sum-root-to-leaf-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//lenght := 0
	// math.Pow10()
	vals := calTree(root)

	sum := 0
	for _, val := range vals {
		//  遍历真正数据
		for j, v := range val {
			acc := v * int(math.Pow10(j))
			fmt.Printf("v: %v \n", acc)
			sum = sum + acc
			// 到了最后，补上跟节点
			// if j == len(val) - 1 {
			// sum = sum + root.Val * int(math.Pow10(j + 1))
			// }
		}
	}
	return sum

}


func calTree(node *TreeNode) [][]int {

	var result [][]int
	if node == nil {
		return result
	}
	// 往下填入数值
	leftNums := calTree(node.Left)

	//往下填入数值
	rightNums :=  calTree(node.Right)
	// 叶子结点
	if len(leftNums) == 0 && len(rightNums) == 0 {
		return [][]int{{node.Val}}
	}
	for _, s := range append(leftNums, rightNums...) {
		// 回溯后加上本次的值
		result = append(result, append(s, node.Val))
	}
	return result
}

// func copy(s []int) []int {
//     n := make([]int, len(s), len(s))
//     copy(n, s)
//     return s
// }