package awesome

func minDistance(word1 string, word2 string) int {


	//因此，我们就可以使用动态规划来解决这个问题了。
	//我们用 D[i][j] 表示 A 的前 i 个字母和 B 的前 j 个字母之间的编辑距离。

	// D[i][j] = min(D[i][j], D[i-1][j] + 1, D[i][j-1] + 1)

	word1Length := len(word1)
	word2Length := len(word2)

	dp := make([][]int, word1Length + 1, word1Length + 1)

	for i := range dp {
		dp[i] = make([]int, word2Length + 1)
	}


	// 初始化
	for i := 0; i <= word1Length; i ++ {
		dp[i][0] = i
	}
	for j := 0; j <= word2Length; j ++ {
		dp[0][j] = j
	}

	for i := 1; i <= word1Length; i ++ {

		for j := 1; j <= word2Length; j ++ {

			x := dp[i-1][j-1]
			if word1[i - 1] != word2[j - 1] {
				x = x + 1
			}
			dp[i][j] = minForThree(x, dp[i-1][j] + 1, dp[i][j-1]+ 1)
		}
		//fmt.Printf("dp: %+v \n", dp)
	}
	return dp[word1Length][word2Length]
}


func minForThree(x, y, z int) int {
	if min(x, y) > z {
		return z
	}
	return min(x, y)
}