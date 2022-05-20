package awesome

//给定一个二叉树的 root ，确定它是否是一个 完全二叉树 。
//
//在一个 完全二叉树 中，除了最后一个关卡外，所有关卡都是完全被填满的，并且最后一个关卡中的所有节点都是尽可能靠左的。它可以包含 1 到 2h 节点之间的最后一级 h 。

func isCompleteTree(root *TreeNode) bool {

	if root == nil {
		return true
	}
	// 不完全二叉树
	if root.Right != nil && root.Left == nil {
		return false
	}
	nodeQueue := make([]*TreeNode, 0, 0)
	curPos := 0
	nodeQueue = append(nodeQueue, root)

	isStartNull := false
	isNotCompleteTree := false

	for ; ; {

		if curPos > len(nodeQueue)-1 {
			break
		}
		curNode := nodeQueue[curPos]
		if curNode != nil && isStartNull {
			isNotCompleteTree = true
			break
		}
		if curNode == nil {
			isStartNull = true
			curPos++
			continue
		}
		nodeQueue = append(nodeQueue, curNode.Left)
		nodeQueue = append(nodeQueue, curNode.Right)
		curPos++
	}

	return !isNotCompleteTree
	// 需要左右都是完全二叉树
	//return isCompleteTree(root.Right) && isCompleteTree(root.Left)
}
