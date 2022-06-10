package awesome

func subsets(nums []int) [][]int {

	//empty := []int{}
	//nums = append(nums, empty)
	// -11 当作空
	//nums = append([]int{-11}, nums...)
	//return reSubsets(nums)
	var set []int
	var ans [][]int
	// 是否选择这个内容
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int{}, set...))
			return
		}
		// 加这个，然后递归选择下一个位置
		set = append(set, nums[cur])
		dfs(cur + 1)
		// 不加这个，然后递归看下一个位置
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return ans
}

func reSubsets(nums []int) [][]int {

	// 递归到极限了
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	var result [][]int
	// 	选择性pick
	for i, pick := range nums {
		sets := reSubsets(nums[i+1:])
		for _, s := range sets {
			if pick == -11 {
				result = append(result, s)
			} else {
				result = append(result, append(s, pick))
			}
		}
	}
	return result
}
