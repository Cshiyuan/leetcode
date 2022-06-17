package awesome

//给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。

func countBits(n int) []int {

	if n == 0 {
		return []int{0}
	}

	dpStep := make([]int, n+1, n+1)

	// 初始化
	dpStep[0] = 0
	dpStep[1] = 1

	//pow := 1

	// 遍历值
	for val := 2; val <= n; val++ {

		yihuoval := val & (val - 1)
		if yihuoval != 0 {
			dpStep[val] = dpStep[yihuoval] + dpStep[val-yihuoval]
		} else {
			dpStep[val] = 1
			//pow ++
		}
	}
	return dpStep

}
