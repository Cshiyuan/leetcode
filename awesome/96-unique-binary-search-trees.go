package awesome

//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的
//二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

func numTrees(n int) int {

	if n == 1 {
		return 1
	}

	// 组件节点组合
	var nodes []int
	for node := 1; node <= n; node++ {
		nodes = append(nodes, node)
	}
	return reNumTrees(nodes)

}

func reNumTrees(nodes []int) int {

	if len(nodes) <= 1 {
		return 1
	}
	if len(nodes) == 2 {
		return 2
	}
	count := 0
	for i := range nodes {
		// 右边 + 左边
		count = count + reNumTrees(nodes[:i])*reNumTrees(nodes[i+1:])
	}
	return count
}
