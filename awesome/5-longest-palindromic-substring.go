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
	return s[maxI : maxJ+1]
}


// 中心扩展
func longestPalindrome_V2(s string) string {

	center := 0

	maxleft := 0
	maxright := 0

	checkMax := func(a, b int) {
		if b - a > maxright - maxleft {
			maxright = b
			maxleft = a
		}
	}

	centerRangeCheck := func(left, right int) {

		for {
			if left < 0 || right >= len(s) {
				break
			}
			if s[left] != s[right] {
				break
			}
			checkMax(left, right)
			left --
			right ++
		}
	}


	for ;center < len(s) * 2 - 1; center++ {


		//|0|1|2|3|4|5|
		var left, right int
		if center % 2 == 0 {
			oriCenter := center / 2
			left = oriCenter
			right = oriCenter
		} else {
			left = center / 2
			right = center / 2 + 1
		}
		centerRangeCheck(left, right)

	}

	return s[maxleft:maxright + 1]

}