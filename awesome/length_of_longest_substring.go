package awesome

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {

	if len(s) == 1 {
		return 1
	}

	end := 0
	start := 0

	str := ""
	maxLength := 0
	for ; end < len(s); end++ {
		// 已经重复了
		if strings.Contains(str, string(s[end])) {
			length := end - start
			// 说明找到一个更大的
			if length > maxLength {
				maxLength = length
			}
			// 置换遍历窗口
			start++
			str = str[1:]
			end --
			continue
		}
		str = str + string(s[end])
	}

	length := end - start
	// 说明找到一个更大的
	if length > maxLength {
		maxLength = length
	}

	return maxLength

}
