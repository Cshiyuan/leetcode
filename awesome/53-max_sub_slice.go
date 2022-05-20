package awesome

import "math"

//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//子数组 是数组中的一个连续部分。
//
//
//
//示例 1：
//
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	// dp初始化
	dpStep := make([]int, len(nums), len(nums))
	dpStep[0] = nums[0]

	for i, n := range nums {
		if i == 0 {
			continue
		}
		dpStep[i] = int(math.Max(float64(dpStep[i-1] + n), float64(n)))
	}

	maxNum := dpStep[0]

	for _ , v := range dpStep {
		if maxNum < v {
			maxNum = v
		}

	}
	return maxNum
}

func maxSubArray2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	leftpos := 0

	for ; ; {

		if leftpos > len(nums)-1 {
			break
		}
		rightpos := len(nums) - 1
		tempSump := math.MinInt64
		for ; rightpos >= leftpos; {
			tempSum := 0
			for _, n := range nums[leftpos : rightpos+1] {
				tempSum = tempSum + n
			}

			if tempSump < tempSum {
				tempSump = tempSum
			}
			rightpos--
		}

		if maxSum < tempSump {
			maxSum = tempSump
		}
		leftpos++

	}
	return maxSum
}

// 递归
func lMaxSubArray(nums []int) (int, int, int) {

	if len(nums) == 1 {
		return nums[0], 0, 0
	}
	// 求后面几个的最大连续数组
	tVal := nums[0]
	// 分而治之
	subList := nums[1:]

	subMaxVal, start, end := lMaxSubArray(subList)
	// 能够直接衔接上
	if start == 0 {
		// 大于0，直接相加
		if tVal > 0 && subMaxVal > 0 {
			return subMaxVal + tVal, 0, end + 1
		}
		if tVal > 0 && subMaxVal < 0 {
			return tVal, 0, 0
		}
		if tVal < 0 {

			if tVal > subMaxVal {
				return tVal, 0, 0
			}
			return subMaxVal, start + 1, end + 1
		}

	}

	if subMaxVal < 0 && tVal > 0 {
		return tVal, 0, 0
	}
	if subMaxVal < 0 && tVal < 0 {
		if tVal > subMaxVal {
			return tVal, 0, 0
		} else {
			return subMaxVal, start + 1, end + 1
		}

	}
	if tVal < 0 && subMaxVal > 0 {
		return subMaxVal, start + 1, end + 1
	}

	// 不能够直接衔接上, 遍历找答案

	tEndIndex := len(subList)

	rEndMaxIndex := tEndIndex
	rSubMaxVal := tVal

	for tEndIndex >= 0 {

		ttSubMaxVal := tVal

		tSubList := subList[:tEndIndex]
		for _, v := range tSubList {
			ttSubMaxVal = ttSubMaxVal + v
		}
		if ttSubMaxVal > rSubMaxVal {
			rSubMaxVal = ttSubMaxVal
			rEndMaxIndex = tEndIndex
		}

		tEndIndex--
	}

	if rSubMaxVal > subMaxVal {
		return rSubMaxVal, 0, rEndMaxIndex + 1
	}

	return subMaxVal, start + 1, end + 1
}
