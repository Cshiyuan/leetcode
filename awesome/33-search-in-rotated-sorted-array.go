package awesome


// 尝试log(N)
func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	centerPos := len(nums) / 2
	if target == nums[centerPos] {
		return centerPos
	}

	// 左边单调有序
	if nums[centerPos] > nums[len(nums)-1] {

		// 说明在左边
		if target < nums[centerPos] && target >= nums[0] {
			return search(nums[:centerPos], target)
		}
		pos := search(nums[centerPos:], target)
		if pos == - 1 {
			return -1
		}
		return centerPos + pos

	} else { //右边单调有序

		var pos int
		// 说明在右边
		if target > nums[centerPos] && target <= nums[len(nums) - 1]{
			pos = search(nums[centerPos:], target)
			if pos == - 1 {
				return -1
			}
			return centerPos  + pos
		}
		return search(nums[:centerPos], target)
	}
}
