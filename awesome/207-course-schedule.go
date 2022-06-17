package awesome

import "fmt"

//你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
//在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
//
//例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/course-schedule
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func canFinish(numCourses int, prerequisites [][]int) bool {

	// 判断环


	// var courses []int
	// courses = make([]int , numCourses, numCourses)
	// courses := make(map[int][]int, 0)
	// // 初始化树
	// for i := 0 ;i < numCourses ; i++{
	//     _, ok := courses[-1]
	//     if !ok {
	//         courses = make([]int, 0, 0)
	//     }
	//     courses[-1] = append(courses[-1], i)
	// }
	edges := make(map[int][]int, 0)
	// 改变指项
	for _, relation := range prerequisites {

		_, ok := edges[relation[1]]
		if !ok {
			edges[relation[1]] = make([]int, 0, 0)
		}
		edges[relation[1]] = append(edges[relation[1]], relation[0])
		// 加入边
		fmt.Printf("relation %v -> %v \n", relation[1], relation[0])
	}

	// 0 未到达， 1，正在搜索 2 搜索结束
	vist := make([]int , numCourses, numCourses)


	isvalid := true
	var dfs func(i int)
	dfs = func(i int) {
		if vist[i] == 2 {
			return
		}
		// 说明数据有问题
		if vist[i] == 1 {
			fmt.Printf("i: %v, vist 1\n", i)
			isvalid = false
			return
		}
		if vist[i] == 0 {
			vist[i] = 1
		}
		edge, ok := edges[i]
		if ok {
			for _, e := range edge {
				dfs(e)  //深度优先
			}
		}
		vist[i] = 2
	}
	for i:= 0; i < numCourses; i ++ {
		dfs(i)
	}
	return isvalid
}