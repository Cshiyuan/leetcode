package awesome

//给定一个 m x n 二维字符网格 board
//和一个字符串单词 word 。如果 word
//存在于网格中，返回 true ；否则，返回 false 。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，
//其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
//同一个单元格内的字母不允许被重复使用。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/word-search
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type Point struct {
	x int
	y int
}

func exist(board [][]byte, word string) bool {

	if len(word) == 0 {
		return true
	}

	var path []*Point
	// 备选集合
	var points []*Point

	// 寻找下一个位置 位置初始化
	points = findStartNextPos(byte(word[0]), board)
	// 没有找到下一个位置
	if len(points) == 0 {
		return false
	}

	isFind := rangePoints(points, word[1:], board, path)

	return isFind
}

func rangePoints(sets []*Point, word string, board [][]byte, path []*Point) bool {

	if len(word) == 0 {
		return true
	}

	// 遍历不同出发点
	for _, p := range sets {

		// 加入路径中
		path = append(path, p)
		// 找集合
		points := findNextPos(p, byte(word[0]), board, path)
		// 没有路径了
		if len(points) == 0 {
			// 去掉路径
			path = path[:len(path)-1]
			continue
		}
		//进入下一个循环
		if rangePoints(points, word[1:], board, path) {
			return true
		}
		// 去掉路径
		path = path[:len(path)-1]
	}

	return false
}

func findNextPos(start *Point, target byte, board [][]byte, path []*Point) []*Point {

	var result []*Point
	// 遍历位置周围
	if start.y+1 < len(board[0]) &&
		board[start.x][start.y+1] == target &&
		!isExit(Point{x: start.x, y: start.y + 1}, path){
		result = append(result, &Point{
			x: start.x,
			y: start.y + 1,
		})
	}

	if start.y-1 >= 0 &&
		board[start.x][start.y-1] == target &&
		!isExit(Point{x: start.x, y: start.y - 1}, path){
		result = append(result, &Point{
			x: start.x,
			y: start.y - 1,
		})
	}

	if start.x-1 >= 0 &&
		board[start.x-1][start.y] == target &&
		!isExit(Point{x: start.x - 1, y: start.y}, path){
		result = append(result, &Point{
			x: start.x - 1,
			y: start.y,
		})
	}

	if start.x+1 < len(board) &&
		board[start.x+1][start.y] == target &&
		!isExit(Point{x: start.x + 1, y: start.y}, path) {
		result = append(result, &Point{
			x: start.x + 1,
			y: start.y,
		})
	}

	return result

}

func findStartNextPos(s byte, board [][]byte) []*Point {

	var result []*Point
	for i, v := range board {
		for j := range v {
			// 元素存在
			if s == board[i][j]  {
				result = append(result, &Point{x: i, y: j})
			}
		}
	}
	return result
}

// 判断是否已经在选择的路径上
func isExit(cur Point, path []*Point) bool {
	for _, s := range path {
		if cur.x == s.x &&
			cur.y == s.y {
			return true
		}
	}
	return false
}
