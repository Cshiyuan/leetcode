package awesome

import "fmt"

//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/house-robber
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func rob(nums []int) int {

	// 动态规划
	// dp[i] = max(dp[i-1], dp[i-1]+ nums[i])
	n := len(nums)
	dp := make([]int, n+1, n+1)
	dp[0] = 0
	dp[1] = nums[0]

	for i:= 2; i <= n; i ++ {
		dp[i] = max(dp[i-1], dp[i-2] + nums[i-1])
		fmt.Printf("i: %v, val: %v\n", i, dp[i])
	}
	return dp[n]
}