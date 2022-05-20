package awesome

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//

var leftv = map[string]bool{
	"(": true,
	"{": true,
	"[": true,
}

var rightv = map[string]bool{
	")": true,
	"{": true,
	"[": true,
}

var vv = map[string]string{
	//'('，')'，'{'，'}'，'['，']'
	"(": ")",
	"{": "}",
	"[": "]",
}

func isValid(s string) bool {

	if len(s) == 1 {
		return false
	}

	stack := make([]string, 0, 0)
	for _, n := range s {
		// 左符号，入栈
		_, ok := leftv[string(n)]
		if ok {
			stack = append(stack, string(n))
			continue
		}
		if len(stack) == 0 {
			return false
		}
		// 不是左就是右
		v := stack[len(stack)-1]
		if vv[v] != string(n) {
			return false
		}


		//出栈
		stack = stack[:len(stack)-1]
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func isValid2(s string) bool {

	if len(s) == 1 {
		return false
	}

	leftPos := 0
	rightPos := len(s) - 1

	for ; leftPos <= rightPos; {

		if vv[string(s[leftPos])] != string(s[rightPos]) {
			return false
		}
		leftPos++
		rightPos--
	}
	return true
}
