package awesome

import "fmt"

func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2 * n - 1 ; i ++ {

		// 长度
		right := 0
		left := 0
		if (i % 2) == 0 {
			right = i / 2
			left = i / 2
		} else {
			right = i / 2
			left = i / 2 + 1
		}

		for ;right >= 0 && left < n; {
			if s[right] == s[left] {
				fmt.Printf("right: %v, left: %v, i: %v \n", right, left, i)
				ans ++
				right --
				left ++
				continue
			}
			break
		}
	}
	return ans
}