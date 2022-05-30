package awesome

//给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findDisappearedNumbers(nums []int) []int {

	n := len(nums)
	sets := make([]int, n, n)
	if n == 0 {
		return []int{}
	}

	for _, val := range nums {
		if val <= n {
			sets[val-1] = 1
		}
	}
	var result []int
	for i, val := range sets {

		if val == 0 {
			result = append(result, i+1)
		}
	}

	return result
}
