package awesome

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
//此外，你可以假设该网格的四条边均被水包围。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/number-of-islands
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func numIslands(grid [][]byte) int {

	cnt := 0
	iLength := len(grid)
	jLength := 0
	if iLength > 0 {
		jLength = len(grid[0])
	}
	var reSet func(i, j int)
	reSet = func(i, j int)  {
		// 置换0
		grid[i][j] = '0'
		if i - 1 >= 0 && grid[i-1][j] == '1' {
			reSet(i-1, j)
		}
		if i + 1 < iLength  && grid[i+1][j] == '1' {
			reSet(i+1,j)
		}
		if j - 1 >= 0 && grid[i][j-1] == '1' {
			reSet(i,j-1)
		}
		if j + 1 < jLength  && grid[i][j+1] == '1' {
			reSet(i,j+1)
		}
	}

	//printGrid := func() {
	//   for _, v := range grid {
	//        fmt.Printf("%v\n", v)

	//   }
	//   fmt.Printf("--------\n")
	// }

	//printGrid()

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				cnt++
				reSet(i, j)
				//printGrid()
			}
		}
	}
	return cnt
}
