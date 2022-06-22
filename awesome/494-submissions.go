package awesome

func findTargetSumWays(nums []int, target int) int {


	// 遍历，回溯


	var ans int

	var dfs func(nums []int, target int)

	dfs = func(nums []int, target int) {

		if len(nums) == 0 {
			if target != 0 {
				return
			} else {
				//   找到一个解
				ans++
				return
			}
		}

		newTarget := target + nums[len(nums) - 1]
		dfs(nums[:len(nums)-1], newTarget)

		newTarget = target - nums[len(nums) -1]
		dfs(nums[:len(nums)-1], newTarget)
	}

	dfs(nums, target)

	return ans
}

// 动态规划
func findTargetSumWays2(nums []int, target int) int {


	// 遍历，回溯

	sum := 0
	for _, v := range nums {
		sum += v
	}

	diff := sum -target
	// 必须非负偶数
	if diff < 0 || diff % 2 == 1 {
		return 0
	}
	n, neg := len(nums), diff/2

	// i 前i个数， j目标值
	dp := make([][]int , n + 1)
	for i := range dp {
		dp[i] = make([]int , neg + 1)
	}

	dp[0][0] = 1
	for i, num := range nums {
		for j := 0; j <= neg; j ++ {
			dp[i+1][j] = dp[i][j]
			if j >= num {  // 可以选
				dp[i+1][j] = dp[i][j] + dp[i][j-num]
			} else {   // 不可以选
				dp[i+1][j] = dp[i][j]
			}
		}
	}



	return dp[n][neg]
}