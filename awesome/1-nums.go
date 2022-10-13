package awesome

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。
func twoSum2(nums []int, target int) []int {

	// 记录可以组成的值
	_, vals := rTwoSum(nums, target)
	// 记录位置
	var posList []int
	sets := make(map[int]bool,  0)
	for _, v := range vals {
		var pos int
		copyNums := make([]int, len(nums), len(nums))
		copy(copyNums, nums)
		pos = findPos(v, copyNums, sets)
		posList  = append(posList, pos)

	}
	return posList
}



func twoSum(nums []int, target int) []int {

	valSet := make(map[int]int, 0)
	rTarget := target
	for i, num := range nums {
		// 记录需要的数
		rTarget = target - num
		valSet[rTarget] = i
	}

	for i, num := range nums {
		v, ok := valSet[num]
		if ok && i != v{
			return []int{i,v}
		}
	}
	return []int{}
}



func findPos(val int, nums []int, sets map[int]bool) (int) {
	for i, n := range nums {
		if val == n {
			_, ok := sets[i]
			if ok {
				continue
			}
			nums = append(nums[:i], nums[i+1:]...)
			sets[i] = true
			return i
		}
	}
	panic("exception!")
}

// 返回的是具体的值
func rTwoSum(nums []int, target int) (bool, []int) {
	if len(nums) == 0 {
		return false, nil
	}

	// 跪了target为0，标准答案不给出空数组
	if target == 0 {
		isHaveEq := false
		for _, n := range nums {
			if n == 0 {
				isHaveEq = true
			}
		}
		if !isHaveEq {
			return true, []int{}
		}
	}

	// 已经递归到最后一步了
	if len(nums) == 1 {
		if nums[0] != target {
			return false, nil
		}
		return true, []int{nums[0]}
	}



	// 尝试通过dp方式去解
	for i, num := range nums {

		//if target - num < 0 {
		//	continue
		//}
		//if target == num {
		//	return true, []int{num}
		//}
		var copyNums = make([]int, len(nums), len(nums))
		copy(copyNums, nums)
		pre := copyNums[:i]
		afr := copyNums[i+1:]
		newNum := append(pre, afr...)

		ok, result := rTwoSum(newNum, target-num)
		// 只要找到一个结果就可以返回了
		if ok {
			return true, append(result, num)
		}
	}
	return false, nil
}
