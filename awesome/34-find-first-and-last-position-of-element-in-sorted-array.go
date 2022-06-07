package awesome

func searchRange(nums []int, target int) []int {

	// 二分法
	// 获得中间位置
	centerPos := len(nums) / 2

	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 && nums[0] != target {
		return []int{-1, -1}
	}
	if len(nums) == 1 && nums[0] == target {
		return []int{0, 0}
	}

	// 说明命中找到了范围
	if nums[centerPos] == target {
		// 合并后的范围
		mergeRange := []int{centerPos, centerPos}
		lRange := searchRange(nums[:centerPos], target)
		// 范围扩展
		if lRange[0] != - 1 {
			mergeRange[0] = lRange[0]
		}
		rRange := searchRange(nums[centerPos:], target)
		// 右边界
		if rRange[1] != -1 {
			mergeRange[1] = rRange[1] + centerPos
		}
		return mergeRange
	}

	// 说明范围在左边
	if target < nums[centerPos] {
		return searchRange(nums[:centerPos], target)
	} else { // 说明范围在右边
		ranges := searchRange(nums[centerPos:], target)
		if ranges[0] == -1 {
			return ranges
		}
		// 补充位置
		ranges[0] = centerPos + ranges[0]
		ranges[1] = centerPos + ranges[1]
		return ranges
	}
}
