package awesome

import "math"

// 动态规划
func numSquares(n int) int {
	f := make([]int, n +1)
	f[0] = 0
	for i := 1; i <= n; i ++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i ; j ++ {
			// 如果刚好能整减，就是f[0] = 0
			minn = min(minn, f[i - j * j])
		}
		f[i] = minn + 1
	}
	return f[n]
}