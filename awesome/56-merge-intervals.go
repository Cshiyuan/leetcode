package awesome

import "sort"

//以数组 intervals 表示若干个区间的集合，其中单个区间为
//intervals[i] = [starti, endi] 。
//请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

func merge(intervals [][]int) [][]int {

	if len(intervals) <= 1 {
		return intervals
	}

	// 按照左端排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 开始合并
	left := intervals[0][0]
	right := intervals[0][1]
	// 存储结果
	var result [][]int
	// 排序完毕
	for i, v := range intervals[1:] {

		if v[0] <= right && v[1] >= right {
			right = v[1]
		}

		if v[0] > right {
			result = append(result, []int{left, right})
			left = v[0]
			right = v[1]
		}
		if i == len(intervals) - 2 {
			result = append(result, []int{left, right})
		}


	}
	return result
}

func merge2(intervals [][]int) [][]int {

	//mergePos := []int{}
	//for i := 0; i <

	return nil
}

func reMerge(b1, b2 []int) (bool, []int) {
	can := func(b1, b2 []int) bool {
		if b1[1] >= b2[0] && b1[0] <= b2[1] {
			return true
		}
		return false
	}
	if can(b1, b2) {
		return true, []int{b1[0], b2[1]}
	}
	if can(b2, b1) {
		return true, []int{b2[0], b1[1]}
	}
	return false, nil
}

// 合并，并且返回新的结合，并告知有没有发生合并
func mergeFirstRange(intervals [][]int) ([][]int, bool) {

	//isMerge
	restPos := []int{}
	mergeIntervals := intervals[0]
	for i, v := range intervals[1:] {
		ok, newMergeInterval := reMerge(mergeIntervals, v)
		// 发生了合并
		if ok {
			mergeIntervals = newMergeInterval
			continue
		}
		restPos = append(restPos, i)
	}

	return nil, false
}
