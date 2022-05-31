package awesome

import "math"

//给定一个长度为 n 的整数数组 height。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
//返回容器可以储存的最大水量。
//
//说明：你不能倾斜容器。
//-;
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/container-with-most-water
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// double pointer
func maxArea(height []int) int {

	i := 0
	j := len(height) - 1

	maxSize := 0
	for {
		if j <= i {
			break
		}
		tHeight := math.Min(float64(height[j]), float64(height[i]))
		size := int(tHeight * float64(j-i))
		// 选择更大值
		if size > maxSize {
			maxSize = size
		}
		if height[j] < height[i] {
			j--
		} else {
			i++
		}
	}

	return maxSize
}

// for - for method
func maxArea2(height []int) int {

	if len(height) < 2 {
		return 0
	}

	// 第一个元素是位置，第二个元素是长度
	// 第一个元素的位置，表示的是左端点
	// 全部初始化为0
	dpSteps := make([][]int, len(height), len(height))
	for i := range dpSteps {
		dpSteps[i] = make([]int, len(height), len(height))
	}

	maxSize := 0

	// 遍历计算左端点位置
	for i := range height {
		// 从长度1开始
		for l := 1; l < len(height)-i; l++ {
			j := i + l
			tHeight := math.Min(float64(height[j]), float64(height[i]))
			size := tHeight * float64(l)
			dpSteps[i][l] = int(math.Max(float64(dpSteps[i][l-1]), size))
			// 选择更大值
			if dpSteps[i][l] > maxSize {
				maxSize = dpSteps[i][l]
			}
		}
	}

	return maxSize
}
