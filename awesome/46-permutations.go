package awesome

func permute(nums []int) [][]int {

	if len(nums) == 0 {
		return [][]int{}
	}

	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	permutation := make([][]int, 0, 0)
	// 第一个选择的元素
	for i, pick := range nums {

		copyNum := make([]int, len(nums), len(nums))
		copy(copyNum, nums)

		// 剩余元素全排列
		e := permute(append(copyNum[:i],copyNum[i+1:]...))
		// 拼接补充
		for _, element := range e {
			permutation = append(permutation, append(element, pick))
		}
	}
	return permutation
}
