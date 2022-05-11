package awesome

type intSlice []int

func quickSort(sortArray intSlice, left int, right int) {

	if left < right {
		key := sortArray[(left+right)/2]
		i := left
		j := right

		for {
			for sortArray[i] < key {
				i++
			}
			for sortArray[j] > key {
				j--
			}
			if i >= j {
				break
			}
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}

		quickSort(sortArray, left, i-1)
		quickSort(sortArray, j+1, right)
	}

}

// 交换方法
func (numbers intSlice) swap(i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}
