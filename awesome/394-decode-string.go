package awesome

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var stk []string
	prt := 0
	for prt < len(s) {

		cur := s[prt]
		// 当前字符
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &prt)
			stk = append(stk, digits)
			// 入栈
		} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' || cur == '[' {
			stk = append(stk, string(cur))
			prt++
			// 需要出栈
		} else if cur == ']'{
			prt ++
			sub := []string{}
			// 一直出栈
			for stk[len(stk) - 1] != "[" {
				sub = append(sub, stk[len(stk) - 1])
				stk = stk[:len(stk) - 1]
			}
			// 反转字符串
			for i := 0; i < len(sub) / 2; i++ {
				sub[i], sub[len(sub) - i - 1] = sub[len(sub) - i - 1], sub[i]
			}// 扩号出栈
			stk = stk[:len(stk) -1]
			// 转数字
			repTime, _ := strconv.Atoi(stk[len(stk) - 1])
			stk = stk[:len(stk) - 1]
			t := strings.Repeat(getString(sub), repTime)
			// 将重复后的放回栈中
			stk = append(stk, t)
		}
	}

	return getString(stk)
}


func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}