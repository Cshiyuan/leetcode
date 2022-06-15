package awesome

import "math"

func maxProduct(nums []int) int {


	// 动态规划 最优子结构
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	maxSum := math.MinInt32
	// 初始化矩阵, j = 长度, i = 起点
	for i := range nums {
		// maxSteps[i][1] = nums[i]
		maxSum = max(nums[i], maxSum)
		// minSteps[i][1] = nums[i]
	}

	// for i := 0; i < n; i ++ {
	maxSteps := make([]int,n+1, n+1)
	minSteps := make([]int,n+1, n+1)
	maxSteps[1] = nums[0]
	minSteps[1] = nums[0]
	// 长度
	for j := 2; (j) - 1 < n; j ++ {
		maxSteps[j] = max(maxSteps[j - 1] * nums[j-1], nums[j-1])
		minSteps[j] = min(minSteps[j-1] * nums[j-1], nums[j-1])
		maxSteps[j] = max(maxSteps[j], minSteps[j-1] * nums[j-1])
		minSteps[j] = min(minSteps[j], maxSteps[j-1] * nums[j-1])
		//fmt.Printf("i: %v, j: %v, val: %v %v \n", i, j, maxSteps[j], minSteps[j])
		// 取更大值
		maxSum = max(maxSteps[j], maxSum)
	}
	// }
	return maxSum
}
