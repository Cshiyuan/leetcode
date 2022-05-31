package awesome

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/letter-combinations-of-a-phone-number
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

var numbers = map[string][]string{
	"1": {""},
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {

	var result []string
	for _, s := range digits {

		if len(result) == 0 {
			result = numbers[string(s)]
			continue
		}

		result = combine(result, numbers[string(s)])
	}
	return result
}

func combine(val []string, val2 []string) []string {

	var result []string
	for _, val := range val {
		result = append(result, newSlice(val, val2)...)
	}
	return result
}

func newSlice(val string, val2 []string) []string {

	var result []string
	for _, v := range val2 {
		result = append(result, val+v)
	}
	return result
}
