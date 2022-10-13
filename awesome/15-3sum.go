package awesome

import (
	"sort"
	"strconv"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}


	var resultList [][]int

	setList := make([]int, 0, len(nums))
	mapNums := make(map[int]int, len(nums))

	zeroCount := 0

	// 1。 数据去重
	for _, n := range nums {

		if n == 0 {
			zeroCount ++
		}

		_, ok := mapNums[n]
		if !ok {
			mapNums[n] = 1
			setList = append(setList, n)
			continue
		}
		mapNums[n] = mapNums[n] + 1
	}
	// 2。 遍历枚举
	for len(setList) > 0 {
		// 基准值
		baseNum := setList[0]
		// 截断, 删除
		setList = setList[1:]


		isOnceMap := make(map[int]bool, 0)


		for _, v := range setList {
			// 	去重
			_, ok := isOnceMap[v]
			if ok {
				continue
			}

			// 需要的值
			key := 0 - (baseNum + v)

			// 	去重
			_, ok = isOnceMap[key]
			if ok {
				continue
			}

			// 不存在这个可以用的值
			count, ok := mapNums[key]
			if !ok {
				continue
			}
			// 存在可以用的值
			if key == baseNum || key == v {
				if count > 1 {
					isOnceMap[v] = true
					resultList = append(resultList, []int{baseNum, v, key})
				}
				continue
			} else {
				isOnceMap[v] = true
				resultList = append(resultList, []int{baseNum, v, key})
			}
		}

		//count := mapNums[baseNum]
		//mapNums[baseNum] = count - 1
		//if count - 1 == 0 {
		//	delete(mapNums, baseNum)
		//}

	}

	if zeroCount >= 3 {
		resultList = append(resultList, []int{0,0,0})
	}


	return repeatRemove(resultList)
}


// 先排序了
func threeSum2(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := - nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first + 1 && nums[second] == nums[second - 1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func repeatRemove(list [][]int ) [][]int {

	setMap := make(map[string]bool, 0)

	resultList := make([][]int, 0 ,0 )

	for _, l := range list {


		sort.Slice(l, func(i, j int) bool {
			return l[i] < l[j]
		})
		str := ""
		for _, v := range l {
			if str != "" {
				str = str + "_"
			}
			str = str + strconv.Itoa(v)
		}

		_, ok := setMap[str]
		if !ok {
			setMap[str] = true
			resultList = append(resultList, l)
		}
	}

	return resultList
}

