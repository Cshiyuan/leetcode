package awesome


//给你一个用字符数组 tasks 表示的 CPU 需要执行的任务列表。其中每个字母表示一种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。在任何一个单位时间，CPU 可以完成一个任务，或者处于待命状态。
//
//然而，两个 相同种类 的任务之间必须有长度为整数 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。
//
//你需要计算完成所有任务所需要的 最短时间 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/task-scheduler
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func leastInterval(tasks []byte, n int) int {

	counter := make(map[byte]int, 0)

	// 统计计数
	for _, t := range tasks {
		counter[t]++
	}

	maxExec, maxExecCnt := 0,0

	for _, c := range counter {

		if c > maxExec {
			maxExec, maxExecCnt = c, 1
		} else if c == maxExec {
			maxExecCnt ++
		}
	}

	if time := (maxExec - 1)*(n+1) +maxExecCnt; time > len(tasks) {
		return time
	}

	return len(tasks)

}