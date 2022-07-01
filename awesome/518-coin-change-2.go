package awesome

func change(amount int, coins []int) int {


	n := len(coins)
	// 记录 dp[i][j] i 表明使用前几个数，j表明使用重量
	dp := make([][]int, n+1, n+1)

	for i := range dp {
		dp[i] = make([]int, amount+1, amount+1)
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {

		for j := 1 ;j <= amount; j++ {

			if j - coins[i-1] >= 0 {

				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {

				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// fmt.Printf("dp: %v \n", dp)

	// return dp[amount]

	return dp[n][amount]
}