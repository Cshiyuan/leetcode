package awesome

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {

	// 可以使用的节点

	target := len(graph) - 1

	var resultPath [][]int
	var runToNext func(node int, path []int)
	//深度遍历
	runToNext = func(node int, path []int){

		nextNodes := graph[node]
		if len(nextNodes) == 0 {
			return
		}
		// 连通的节点
		for _, val := range nextNodes {
			// 找到结果
			if val == target {
				// fmt.Printf("target path: %v \n", path)
				targetPath := append([]int{}, append(path, val)...)
				fmt.Printf("resultPath path: %v \n", targetPath)
				// resultPath = append(resultPath, targetPath)

				continue
			}
			path = append(path, val)
			// fmt.Printf("1 path: %v \n", path)
			runToNext(val, path)
			path = path[:len(path) - 1]
			// fmt.Printf("2 path: %v \n", path)
		}
		return
	}

	paths := []int{0}
	runToNext(0, paths)

	return resultPath
}

