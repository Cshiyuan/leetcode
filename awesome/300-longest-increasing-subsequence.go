package awesome

//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
//子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
//例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

// 动态规划 O n^2
func lengthOfLIS(nums []int) int {

	// 决策

	var ans int
	n := len(nums)
	dp := make([]int, n, n)

	dp[0] = 1
	ans = 1

	for i := 1; i <n ; i ++ {

		dp[i] = 1
		for j := 0; j < i ; j ++ {
			// 无法满足条件
			if nums[j] >= nums[i] {
				continue
			}
			//fmt.Printf("[%v] %v -> %v \n", dp[i], j, i)
			//  取最大值
			dp[i] = max(dp[i], dp[j] + 1)
		}
		ans = max(ans, dp[i])
	}
	return ans
}

// tail规划，二分查找O（N * logN）
func lengthOfLIS2(nums []int) int {

	// 决策
	n := len(nums)

	if n <= 1 {
		return n
	}

	tail := make([]int, n, n)
	// 遍历第一个数，直接放在有序数组tail的开头
	tail[0] = nums[0]
	// 表示有序数组tail的最后一个已赋值元素的索引
	end := 0

	for i := 1; i < n ; i ++ {

		left := 0
		// 遍历当前的数，有可能比有序数组tail实际有效的末尾的那个元素还大
		right := end + 1

		for left < right {
			mid := left + (right - left) / 2
			if tail[mid] < nums[i] {
				// 中位数肯定不是要找的数，把它写在分支前面
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 找到第一个，大于等于nums[i]的元素
		tail[left] = nums[i]
		if left == end + 1 {
			end ++
		}
	}
	end ++
	return end
}
