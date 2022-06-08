package awesome

import (
	"sort"
)

//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
//字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/group-anagrams
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}


func groupAnagrams2(strs []string) [][]string {
	var result [][]string
	for _, s := range strs {
		isInsert := true
		for i, resultStr := range result {
			if isAnagrams(resultStr[0], s) {
				result[i] = append(resultStr, s)
				isInsert = false
				break
			}
		}
		// 说明还没有被归类
		if isInsert {
			result = append(result, []string{s})
		}
	}
	return result
}

func isAnagrams(str1, str2 string) bool {

	// 通过字母计数，来判断相等性
	counterSets1 := make(map[string]int)
	counterSets2 := make(map[string]int)

	for _, s1 := range str1 {
		counterSets1[string(s1)] ++
	}
	for _, s2 := range str2 {
		counterSets2[string(s2)] ++
	}

	for k, v := range counterSets1 {
		v2, ok := counterSets2[k]
		if !ok || v2 != v {
			return false
		}
	}

	for k, v := range counterSets2 {
		v2, ok := counterSets1[k]
		if !ok || v2 != v {
			return false
		}
	}

	return true
}
