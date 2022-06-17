package awesome

import "fmt"

//在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
func maximalSquare(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}

	//  初始化dp
	dp := make([][]int, len(matrix), len(matrix))

	maxSize := 0

	for i := range dp {
		dp[i] = make([]int, len(matrix[0]), len(matrix[0]))
		for j := range dp[i] {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				maxSize = 1
			}
		}
	}

	// 动态规划
	for i := 1; i < len(matrix); i ++ {
		for j := 1; j < len(matrix[0]); j++ {

			if matrix[i][j] != '1' {
				continue
			}
			val := min(dp[i-1][j], dp[i-1][j-1])
			val = min(val, dp[i][j-1])
			dp[i][j] = val + 1
			fmt.Printf("dp[%v][%v]: %v \n", i, j ,dp[i][j])

			maxSize = max(maxSize, dp[i][j])
		}
	}
	return maxSize * maxSize
}