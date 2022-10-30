package awesome

import "sort"

func permuteUnique(nums []int) [][]int {

	// 排序防止出重
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int

	perm := []int{}

	vis := make([]bool, n)

	var backtrack func(int)

	backtrack = func(idx int) {
		if idx == n {
			// 拷贝一份代码
			ans = append(ans, append([]int{}, perm...))
			return
		}

		for i, v := range nums {

			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm) - 1]
		}
		return
	}
	backtrack(0)
	return ans
}