package awesome

import "fmt"

func rotate(matrix [][]int) {

	// 这是一个正方形
	l := len(matrix)
	// 水平翻转
	for i := range matrix {
		if i >= l/ 2 {
			break
		}
		for j := range matrix[i] {
			pos2pos(i, j, l - i - 1, j, matrix)
		}
	}

	// 对角线翻转
	for i := range matrix {
		for j := range matrix[i] {
			if  j > i {
				break
			}
			pos2pos(i, j, j, i, matrix)
			fmt.Println(matrix)
		}
	}

}

func pos2pos(lPos1, lPos2 int, rPos1, rPos2 int, matrix [][]int) {

	temp := matrix[lPos1][lPos2]
	matrix[lPos1][lPos2] = matrix[rPos1][rPos2]
	matrix[rPos1][rPos2] = temp
	fmt.Println(matrix)
}
