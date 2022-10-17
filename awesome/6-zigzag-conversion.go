package awesome

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	counts := make([][]string, numRows, numRows)
	//for i := range counts {
	//	counts[i] = make([]string, len(s), len(s))
	//}

	i := 0
	j := 0
	reRange := false

	for _, s := range s {
		counts[i] = append(counts[i], string(s))
		fmt.Printf("i: %v, j: %v, val: %v \n", i, j, string(s))
		if i == numRows - 1 {
			reRange = true
		}
		if i == 0 {
			reRange = false
		}
		if !reRange {
			i ++
		} else {
			j ++
			i --
		}
	}

	str := ""
	for i := range counts {
		for j := range counts[i] {
			if counts[i][j] != "" {
				str = str + counts[i][j]
			}
		}
	}
	return str
}

// 压缩矩阵
func convert2(s string, numRows int) string {
	r := numRows
	// 特殊情况判断
	if r == 1 || r >= len(s) {
		return s
	}
	// 压缩矩阵
	mat := make([][]byte, r)
	t, x := r*2 -2 , 0
	for i, ch := range s {

		mat[x] = append(mat[x], byte(ch))
		if i%t < r - 1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}