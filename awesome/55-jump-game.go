package awesome

//给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//判断你是否能够到达最后一个下标。

// 标准答案
func canJump2(nums []int) bool {

	l := len(nums)
	rightMost := 0
	for i := 0; i <= rightMost; i++ {
		if rightMost >= l-1 {
			return true
		}
		if nums[i]+i > rightMost {
			rightMost = nums[i] + i
		}
	}
	return false

}

func canJump(nums []int) bool {

	oldRange := 0
	set := []int{0}
	for {
		var newRange int
		set, newRange = getRange(set, oldRange, nums)
		if newRange >= len(nums)-1 {
			return true
		}
		if newRange == oldRange {
			return false
		}
		oldRange = newRange
	}

}

// 输入一个范围，寻找最大的range，并且返回下一步可以到达的下标位置
// 结束条件，newRange > nums.len - 1  true
// newRange <= oldRange false
func getRange(set []int, oldRange int, nums []int) ([]int, int) {

	newRange := oldRange
	// 遍历获得最大范围
	for _, s := range set {
		// 已经超过范围了
		if s > len(nums)-1 {
			newRange = len(nums) - 1
			continue
		}
		numJumpRange := nums[s] + s
		if numJumpRange > newRange {
			newRange = numJumpRange
		}
	}
	// 有扩展范围，但是没有扩展到可以到终点到地步
	if newRange > oldRange && newRange < len(nums) {

		var newSet []int
		for i := oldRange + 1; i <= newRange; i++ {
			newSet = append(newSet, i)
		}
		return newSet, newRange
	}
	return []int{}, newRange
}
