package awesome

import (
	"fmt"
)

//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
//
//注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/word-break
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//https://leetcode.cn/problems/word-break/
func wordBreak(s string, wordDict []string) bool {


	n := len(s)
	if n == 0 {
		return true
	}

	// 通过动态规划来解决
	dp := make([]bool, n+1, n+1)
	// 0 ~ i
	dp[0] = true

	wordDictMap := make(map[string]interface{} , 0)
	for _, word := range wordDict {
		wordDictMap[word] = struct{}{}
	}

	isInDict := func(s string) bool {
		_, ok := wordDictMap[s]
		return ok
	}
	// 动态规划
	for i := 1 ;i <= n; i ++ {
		for j := 0; j < i; j++ {

			if  dp[j] && isInDict(s[j:i]) {
				fmt.Printf("i: %v, j: %v, s: %v", i, j, s[j:i])
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}