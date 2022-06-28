package awesome

import (
	"fmt"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {


	// flights 表明的是航班路线
	// k步，广度优先搜索？

	dpStep := make([][]int, n, n)
	// 初始化矩阵
	for i := range dpStep {
		dpStep[i] = make([]int, n, n)
		for j := range dpStep[i] {
			// 默认最大值
			dpStep[i][j] = math.MaxInt32
		}
		dpStep[i][i] = 0   //到自身的距离都是ok
	}

	// 转换为矩阵
	for i := range flights {
		e := flights[i]
		// src -> dst distance
		dpStep[e[0]][e[1]] = e[2]
	}

	fmt.Printf("dpStep: %v \n", dpStep)


	// f[k][dst] , dst, src订死
	f := make([][]int, k+2, k+2)
	for i := range f {
		f[i] = make([]int, n, n)
		for  j := range f[i] {
			// 到当前位置
			if j == src {
				f[i][j] = 0
			} else {
				f[i][j] = math.MaxInt32
			}
		}
		fmt.Printf("f[%v]: %v \n", i, f[i])

	}
	fmt.Printf("\n")

	for i := 1 ; i <= k + 1; i ++ {

		for p := 0; p < n; p ++ {

			for d := 0; d < n; d++ {

				f[i][p] = min(f[i-1][d] + dpStep[d][p], f[i][p])
			}
		}
		fmt.Printf("f[%v]: %v \n", i, f[i])
	}

	if f[k+1][dst] == math.MaxInt32 {
		return -1
	}

	return f[k+1][dst]
}