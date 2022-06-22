package awesome

import (
	"fmt"
	"math"
)


// 合适的动态规范
func maxProfit23(prices []int) int {

	// 动态规划
	n := len(prices)

	if n <= 1 {
		return 0
	}
	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]


	for i := 1; i < n; i++ {
		// 手上持有股票，那么要不就是前一天没有抛，要不就是今天刚买
		dp[i][0] = max(dp[i-1][0], dp[i-1][2] - prices[i])
		// 手上不持有股票，但在冷冻期 就是今天刚卖掉
		dp[i][1] = dp[i-1][0] + prices[i]
		// 手上不持有股票，不在冷冻期, 前一天在冷冻期
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}

	return max(dp[n-1][1], dp[n-1][2])
}


// 不合适的枚举方式
func maxProfit2(prices []int) int {

	// 动态规划
	n := len(prices)

	if n <= 1 {
		return 0
	}

	dp := make([]int, n , n)
	// dp[i] 为界限
	var ans int

	dp[1] = max(0, prices[1] - prices[0])
	ans = max(ans, dp[1])

	cache := make(map[string]int ,0)


	findMinByCache := func(i, j int) int {
		key := fmt.Sprintf("%v-%v", i, j)
		val, ok := cache[key]
		if ok {
			return val
		}
		val = findMin(prices[i:j])
		cache[key] = val
		return val
	}

	for i := 2; i < n; i++ {

		for j := 0; j < i ; j ++ {

			// 没有任何交易
			dp[i] = max(dp[i], prices[i] -  findMinByCache(0,i))
			// 有交易，但不再交易
			dp[i] = max(dp[i], dp[j])
			/// 看看能不能做交易, 假如做交易
			if i >= j + 2 {
				val := dp[j] + prices[i] - findMinByCache(j+2, i)
				dp[i] = max(dp[i], val)
			}
		}
		fmt.Printf("dp[%v]: %v \n", i, dp[i])
		ans = max(dp[i], ans)
	}

	return ans
}

func findMin(s []int) int {
	m := math.MaxInt32
	for i := range s {
		m = min(m ,s[i])
	}
	return m
}