package awesome

import (
	"sort"
	"strconv"
)

//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//
//计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回-1 。
//
//你可以认为每种硬币的数量是无限的。
//
//
//
//示例1：
//
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//示例 2：
//
//输入：coins = [2], amount = 3
//输出：-1
//示例 3：
//
//输入：coins = [1], amount = 0
//输出：0

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}
	// 动态规划steps
	// 初始化
	dpSteps := make([]int, amount + 1, amount + 1)
	for i, _ := range dpSteps {
		dpSteps[i] = -1
	}
	dpSteps[0] = 0

	// 开始规划
	for n := 1; n <= amount; n++ {

		minValue := -1
		for _, coin := range coins {
			if n < coin {
				//dpSteps[n]
				continue
			}
			// 没有最优解
			if dpSteps[n-coin] == - 1 {
				continue
			}
			if minValue == -1 {
				minValue = dpSteps[n-coin] + 1
				continue
			}
			if minValue > dpSteps[n-coin]+1 {
				minValue = dpSteps[n-coin] + 1
			}
		}
		dpSteps[n] = minValue
	}
	return dpSteps[amount]
}

func coinChange2(coins []int, amount int) int {
	// 排序
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	cache := make(map[int]int, 0)
	return coinChangeInner2(coins, amount, cache)
}

func coinChangeInner2(coins []int, amount int, cache map[int]int) int {

	if amount <= 0 {
		return 0
	}
	sets := make(map[int]interface{}, len(coins))
	for _, coin := range coins {
		sets[coin] = struct{}{}
	}
	// 直接一个就可以cover住
	if _, ok := sets[amount]; ok {
		return 1
	}
	minCount := -1
	// 基本case安排
	for _, coin := range coins {
		if coin > amount {
			continue
		}
		tempAmount := amount - coin
		nextCount, ok := cache[tempAmount]
		if !ok {
			// 通过分而治之的思想
			nextCount = coinChangeInner2(coins, tempAmount, cache)
			cache[tempAmount] = nextCount
		}
		// 说明有匹配结果
		if nextCount != -1 {
			count := 1 + nextCount
			// 没有初始化
			if minCount == -1 {
				minCount = count
			}
			if count < minCount {
				minCount = count
			}
		}
	}
	return minCount
}

func coinChangeInner(coins []int, amount int, record map[string]map[int]int) int {

	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	sets := make(map[int]interface{}, len(coins))
	for _, coin := range coins {
		sets[coin] = struct{}{}
	}

	// 直接一个就可以cover住
	if _, ok := sets[amount]; ok {
		return 1
	}

	//count := 0
	minCount := -1

	// 基本case安排
	for _, coin := range coins {
		if coin > amount {
			continue
		}
		// 统计数量
		count := amount / coin
		for ; ; {

			targetAmount := amount - (count * coin)
			targetCoins := coins[1:]
			key := toStr(targetCoins)
			nextCount, ok := record[key][targetAmount]
			if !ok {
				// 通过分而治之的思想
				nextCount = coinChange(targetCoins, targetAmount)
				if record[key] == nil {
					record[key] = make(map[int]int)
				}
				record[key][targetAmount] = nextCount
			}
			// 说明有匹配结果
			if nextCount != -1 {
				loopCount := count + nextCount
				if minCount == -1 {
					minCount = loopCount
				}
				if loopCount < minCount {
					minCount = loopCount
				}
			}
			count--
			if count == 0 {
				break
			}
		}
	}
	return minCount
}

func toStr(nums []int) string {

	key := ""
	for _, n := range nums {
		if key != "" {
			key = key + "_"
		}
		key = key + strconv.Itoa(n)
	}
	return key
}
