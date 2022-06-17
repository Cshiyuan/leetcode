package awesome

//给你一个含 n 个整数的数组 nums ，
//其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n]
//范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findDisappearedNumbers(nums []int) []int {

	//var result []int

	n := len(nums)

	tVal := n
	for _, val := range nums {
		tVal = tVal ^ val
	}

	return []int{tVal}
}
