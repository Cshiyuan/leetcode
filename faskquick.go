package main


// 升序
func fastQuick(list []int) []int {

	if len(list) == 0 {
		return list
	}
	// 简单取第一个作为基准
	length := len(list)

	low := 0
	high := length - 1
	pivot := list[low]

	for low < high {

		// 都比基准值大
		for low < high && pivot <= list[high] {
			high --
		}

		list[low] = list[high]

		for low < high && pivot >= list[low] {
			low ++
		}
		list[high] = list[low]
	}
	list[high] = pivot
	// 递归
	fastQuick(list[low +1:])
	fastQuick(list[:low])

	return list
}