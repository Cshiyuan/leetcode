package awesome

//给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，
//原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
//我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//必须在不使用库的sort函数的情况下解决这个问题。
//

func sortColors(nums []int) {

	counter := make(map[int]int, 0)

	for _, v := range nums {
		counter[v] ++
	}
	pos := 0
	for _, v := range []int{0, 1, 2} {
		count := counter[v]
		for i := pos; i < pos + count; i++ {
			nums[i] = v
		}
		pos = pos + count
	}
}
