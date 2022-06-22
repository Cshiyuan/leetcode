package awesome

import (
	"fmt"
	"sort"
)

type element struct {
	val int
	count int
}

func topKFrequent(nums []int, k int) []int {

	mapStr := make(map[int]*element, len(nums))

	for _, n := range nums {

		e, ok := mapStr[n]
		if ok {
			e.count ++
			continue
		}
		mapStr[n] = &element{
			val: n,
			count: 1,
		}
	}

	sliceNums := make([]*element, 0, len(nums))
	for _, n := range mapStr {
		sliceNums = append(sliceNums, n)
	}

	fmt.Printf("sliceNums: %v", sliceNums)

	sort.Slice(sliceNums, func(i, j int) bool {
		return sliceNums[i].count > sliceNums[j].count
	})

	var ans []int

	for _, v := range sliceNums[:k] {
		ans = append(ans, v.val)
	}


	return ans
}


func topKFrequent2(nums []int, k int) []int {


	mapStr := make(map[int]int, len(nums))
	for _, n := range nums {
		_, ok := mapStr[n]
		if ok {
			mapStr[n] ++
			continue
		}
		mapStr[n] = 1
	}

	// 频率做下标
	pos := make([][]int, len(nums)+1, len(nums)+1)
	for k, v := range mapStr {
		val := pos[v]
		if val == nil {
			pos[v] = make([]int, 0, 0)
		}
		pos[v] = append(pos[v], k)
	}

	var ans []int
	for i := len(nums); i > 0 && len(ans) < k ; i -- {
		if pos[i] == nil {
			continue
		}
		for _, v := range pos[i] {
			ans = append(ans, v)
		}
	}

	return ans
}