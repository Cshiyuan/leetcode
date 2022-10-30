package awesome

func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var (
		rows, columns = len(matrix), len(matrix[0])
		order = make([]int, rows * columns) //直接预先分配
		index = 0
		left, right, top, bottom = 0, columns - 1, 0, rows - 1
	)
	// 从外往里遍历
	for left <= right && top <= bottom {

		// 只修改列，从左往右
		for column := left; column <= right; column ++ {
			order[index] = matrix[top][column]
			index++
		}
		// 只修改行，从上往下
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index ++
		}
		if left < right && top < bottom {
			// 从右往左
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			// 从下往上
			for row := bottom; row > top; row -- {
				order[index] = matrix[row][left]
				index++
			}
		}
		// 缩小整个圈
		left ++
		right --
		top ++
		bottom --
	}
	return order
}