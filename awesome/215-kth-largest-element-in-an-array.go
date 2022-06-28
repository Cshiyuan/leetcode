package awesome


func QuickSort(list []int, low, high int)  {
	if high > low {
		// 位置划分
		privot := parition(list, low, high)

		// 左边部分排序
		QuickSort(list, low, privot - 1)
		QuickSort(list, privot + 1, high)
	}
	return
}

func parition(list []int, low, high int) int {

	pivot := list[low]

	for low < high {

		for low < high && pivot <= list[high] {
			high --
		}
		//
		list[low] = list[high]

		for low < high && pivot >= list[low] {
			low ++
		}

		list[high] = list[low]
	}
	list[low] = pivot

	return low
}