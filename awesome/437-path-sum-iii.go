package awesome

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//
//路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/path-sum-iii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func pathSum(root *TreeNode, targetSum int) int {


	// dfs

	var ans int
	var dfs func(node *TreeNode, currentSum int)

	// 计算截断位置

	sets := make(map[*TreeNode]bool, 0)
	dfs = func(node *TreeNode, currentSum int) {

		if node == nil {
			//fmt.Printf(" node: null return\n")
			return
		}

		currentSum = currentSum + node.Val
		if currentSum == targetSum {
			//fmt.Printf("node: %v, currentSum: %v\n", node.Val, currentSum)
			ans ++
		}

		// 选择了这个节点，往下试探
		dfs(node.Left, currentSum)
		dfs(node.Right, currentSum)

		// 防止频繁截断尝试
		_, ok := sets[node]
		if !ok {
			//fmt.Printf("cur node: %v, currentSum: %v\n", node.Val, currentSum)
			// 不选择这个节点，截断试探
			dfs(node.Left, 0)
			dfs(node.Right, 0)
			//fmt.Printf("return end %v \n", node.Val)
			sets[node] = true
		}


		return
	}

	dfs(root, 0)

	return ans

}

// 	前缀和解法
func pathSum2(root *TreeNode, targetSum int) int {

	preSum := make(map[int]int)
	preSum[0] = 1

	var dfs func(*TreeNode, int)

	var ans int
	dfs = func(node *TreeNode, curr int) {
		if node == nil {
			return
		}
		curr += node.Val
		// 根据前缀和判断
		ans += preSum[curr-targetSum]

		preSum[curr]++

		dfs(node.Left, curr)
		dfs(node.Right, curr)

		preSum[curr]--
		return

	}
	dfs(root, 0)
	return ans
}


