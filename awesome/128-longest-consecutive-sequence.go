package awesome

// https://leetcode.cn/problems/longest-consecutive-sequence/solution/zui-chang-lian-xu-xu-lie-by-leetcode-solution/
//然后在内层循环中匹配连续序列中的数，因此数组中的每个数只会进入内层循环一次
func longestConsecutive(nums []int) int {

	set := make(map[int]interface{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}

	ans := 0
	for _, v := range nums {
		_, ok := set[v-1]
		if ok {
			continue
		}
		count := 1
		ans = max(ans, count)
		v++
		for ;; {
			_, ok := set[v]
			if !ok {
				break
			}
			count ++
			ans = max(ans, count)
			v++
		}
	}
	return ans
}