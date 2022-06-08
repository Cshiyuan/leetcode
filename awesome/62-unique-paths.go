package awesome

type Point struct {
	m int
	y int
}

func uniquePaths(m int, n int) int {

	// 需要排列组合，右多少种组合

	steps := make(map[Point]int)
	count := 0
	for {
		if m == 1 && n == 0 {
			count++
			break
		}
		if n == 0 && m == 1 {
			count++
			break
		}

		// 往下走 + 往上走 + 1
		count := rePaths(m-1, n) + rePaths(m, n-1) + 1
		return count
	}

}

func rePaths(m, n int) int {

	if m == 1 && n == 0 {
		return 1
	}
	if n == 0 && m == 1 {
		return 1
	}
	// 往下走 + 往上走 + 1
	count := rePaths(m-1, n) + rePaths(m, n-1) + 1
	return count
}
