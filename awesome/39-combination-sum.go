package awesome

//给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
//candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。 
//
//对于给定的输入，保证和为 target 的不同组合数少于 150 个。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/combination-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func combinationSum(candidates []int, target int) [][]int {

	// 初始化dps的中间结果
	steps := make([][][]int, target+1, target+1)

	sets := reCombinationSum(candidates, target, steps)

	return sets

}

// 方便递归回调
func reCombinationSum(candidates []int, target int, steps [][][]int) [][]int {

	val := steps[target]
	// 命中缓存
	if val != nil {
		return val
	}
	for _, c := range candidates {

		// 不能使用
		rest := target - c
		if rest < 0 {
			continue
		}
		if rest == 0 {
			//val := steps[target]
			//if val == nil { //还没初始化
			//	steps[target] = make([][]int, 0, 0)
			//}
			steps[target] = append(steps[target], []int{c})
			continue
		}
		sets := reCombinationSum(candidates, rest, steps)
		// 填充组合
		for _, set := range sets {
			steps[target] = append(steps[target], append(set, c))
		}
	}
	return steps[target]
}

func addSet(steps [][]int, e []int) {

	have := false
	for _, s := range steps {
		if isEqual(s, e) {
			have = true
			break
		}
	}
	if !have {
		steps = append(steps, e)
	}
}

func isEqual(e1 []int, e2 []int) bool {

	sets1 := make(map[int]int, 0)
	sets2 := make(map[int]int, 0)
	for _, e := range e1 {
		sets1[e] ++
	}
	for _, e := range e2 {
		sets2[e] ++
	}
	for k := range sets1 {
		if sets1[k] != sets2[k] {
			return false
		}
	}
	for k := range sets2 {
		if sets1[k] != sets2[k] {
			return false
		}
	}
	return true
}
