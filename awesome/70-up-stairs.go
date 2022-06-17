package awesome

// n = n -1 的组成方式 + 1 + n - 2的组成方式  - 2
// f(n) = f(n-1) + f(n -2)
// 由此可得，通过dp的方式解决
//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
func climbStairs(n int) int {

	dpStep := make([]int, n+2, n+2)

	dpStep[0] = 0
	dpStep[1] = 1
	dpStep[2] = 2

	for i := 3; i <= n; i++ {
		dpStep[i] = dpStep[i-1] + dpStep[i-2]
	}
	return dpStep[n]
}
