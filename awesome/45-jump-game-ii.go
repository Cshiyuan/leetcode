package awesome

func jump(nums []int) int {


	position := len(nums) - 1
	steps := 0
	// 反查位置
	for position > 0 {


		for i := 0; i < position; i++ {
			// 反查距离最远的位置
			if i + nums[i] >= position {
				position = i
				steps ++
				break
			}
		}
	}
	return steps
}