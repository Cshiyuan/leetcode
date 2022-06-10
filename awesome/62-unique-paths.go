package awesome

func uniquePaths(m int, n int) int {

	// 需要的向下，向右的次数
	//goDownCount := m - 1
	//goRightCount := n - 1
	m--
	n--
	steps := make(map[int]map[int]int)
	// 需要排列组合，右多少种组合
	if n == 0 {
		return 1
	}
	if m == 0 {
		return 1
	}
	// 往下走 + 往上走 + 1
	count := rePaths(m-1, n, steps) + rePaths(m, n-1, steps)
	return count

}

func rePaths(m, n int, steps map[int]map[int]int) int {

	if n == 0 {
		return 1
	}
	if m == 0 {
		return 1
	}

	count1, ok := getVal(m-1, n, steps)
	if !ok {
		count1 = rePaths(m-1, n, steps)
	}

	count2, ok := getVal(m, n-1, steps)
	if !ok {
		count2 = rePaths(m, n-1, steps)
	}

	// 往下走 + 往上走 + 1
	count := count1 + count2
	setVal(count, m, n, steps)
	return count
}

func getVal(m, n int, steps map[int]map[int]int) (int, bool) {
	_, ok := steps[m]
	if !ok {
		steps[m] = make(map[int]int)
	}
	val, ok := steps[m][n]
	if !ok {
		return -1, false
	}
	return val, true
}

func setVal(val, m, n int, steps map[int]map[int]int) {
	_, ok := steps[m]
	if !ok {
		steps[m] = make(map[int]int)
	}
	steps[m][n] = val
}

func uniquePathsForDp(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
