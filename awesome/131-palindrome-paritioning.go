package awesome

import "fmt"

func partition(s string) [][]string {

	var ans [][]string
	f := make([][]int8, len(s), len(s))
	for i := range f {
		f[i] = make([]int8, len(s), len(s))
	}

	fmt.Printf("s len: %v", len(s))

	splits := []string{}

	var dfs func(int)
	dfs = func(i int) {
		// 遍历到最后了
		if i == len(s) {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < len(s); j++ {
			if isPalindrome3(s, f, i, j) > 0 {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				// 回退
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return ans
}

//  // 0 表示尚未搜索，1 表示是回文串，-1 表示不是回文串
func isPalindrome3(s string, f [][]int8, i, j int) int8 {
	if i >= j {
		return 1
	}
	if f[i][j] != 0 {
		return f[i][j]
	}
	f[i][j] = -1
	if s[i] == s[j] {
		f[i][j] = isPalindrome3(s, f, i+1, j-1)
	}
	return f[i][j]
}
