package awesome

import "fmt"

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