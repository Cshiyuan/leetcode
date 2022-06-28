package awesome

//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
func canPartition(nums []int) bool {

	// 总数
	sum := 0
	// 最大值
	max := 0
	for _, num := range nums {
		sum = sum + num
		if num > max {
			max = num
		}
	}
	// 不是偶数
	if sum % 2 != 0 {
		return false
	}

	target := sum / 2

	n := len(nums)
	// 数量不足，或者最大值已经超过了需要取的值
	if n < 2 || max > target {
		return false
	}

	// dp[i][j] 0 ~ i 位置，j为目标值
	dp := make([][]bool, n, n)

	for i := range dp {
		dp[i] = make([]bool, target+1,  target+1)
		dp[i][0] =  true

	}
	dp[0][nums[0]] = true

	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j <= target; j ++ {
			// 能放进去
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n-1][target]
}