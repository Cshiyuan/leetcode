package awesome

import "math"

//给定一个包含非负整数的 m x n 网格 grid ，
//请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
//说明：每次只能向下或者向右移动一步。

func minPathSum(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	steps := make([][]int, len(grid), len(grid))
	// 初始化动态规划的dp
	sum := 0
	for i := range steps {
		steps[i] = make([]int, len(grid[0]), len(grid[0]))
		sum = sum + grid[i][0]
		steps[i][0] = sum
	}
	sum = 0
	for j := range grid[0] {
		sum = sum + grid[0][j]
		steps[0][j] = sum
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			steps[i][j] = int(math.Min(float64(steps[i-1][j]),
				float64(steps[i][j-1]))) + grid[i][j]
		}
	}
	return steps[len(grid)-1][len(grid[0])-1]
}
