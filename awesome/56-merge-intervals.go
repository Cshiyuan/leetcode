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
