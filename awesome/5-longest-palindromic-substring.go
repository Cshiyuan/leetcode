package awesome

//给你一个字符串 s，找到 s 中最长的回文子串。
func longestPalindrome(s string) string {

	dpSteps := make([][]bool, len(s), len(s))

	// 初始化动态规范方程
	for i, _ := range dpSteps {
		dpSteps[i] = make([]bool, len(s), len(s))
		dpSteps[i][i] = true
	}

	maxLength := -1
	maxI := 0
	maxJ := 0

	for l := 1; l < len(s); l++ {
		// 开始规划
		for i := 0; i < len(s); i++ {
			j := i + l
			// 已经超出范围
			if j >= len(s) {
				break
			}
			ispalindroma := false
			if s[i] == s[j] {
				ispalindroma = true
			}
			if l <= 2 {
				if ispalindroma {
					if l > maxLength {
						maxLength = l
						maxI = i
						maxJ = j
					}
					dpSteps[i][j] = true
					continue
				}
				dpSteps[i][j] = false
				continue
			}
			dpSteps[i][j] = ispalindroma && dpSteps[i+1][j-1]
			if dpSteps[i][j] == true {
				if l > maxLength {
					maxLength = l
					maxI = i
					maxJ = j
				}
			}
		}
	}

	//	for j := 0; j < len(s); j++ {
	//		if i == j {
	//			continue
	//		}
	//		ispalindroma := false
	//		if s[i] == s[j] {
	//			ispalindroma = true
	//		}
	//		dpSteps[i][j] = ispalindroma && dpSteps[i][j-1]
	//	}
	//}

	return s[maxI : maxJ+1]
}
