package awesome

//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
func rightSideView(root *TreeNode) []int {

	var rightSideView []int

	// 返回深度
	var deepRange func(node *TreeNode, length int) int
	deepRange = func(node *TreeNode, length int) int{

		if node == nil {
			return length
		}
		if length == 0 {
			rightSideView = append(rightSideView, node.Val)
		} else {
			length --
		}
		rightHeight := deepRange(node.Right, length)
		leftHeight := deepRange(node.Left, rightHeight)
		return max(rightHeight, max(length, leftHeight)) + 1
	}
	deepRange(root, 0)
	return rightSideView
}

// 用了map存放
func rightSideView2(root *TreeNode) []int {

	var rightSideView []int

	set := make(map[int]int8, 0)
	addRightSideView := func(val int, length int) {
		_, ok := set[length]
		if !ok {
			rightSideView = append(rightSideView, val)
			set[length] = 1
		}
	}

	// 返回深度
	var deepRange func(node *TreeNode, length int)
	deepRange = func(node *TreeNode, length int) {

		if node == nil {
			return
		}
		addRightSideView(node.Val, length)
		deepRange(node.Right, length + 1)
		deepRange(node.Left, length + 1)
		return
	}
	deepRange(root, 0)
	return rightSideView
}
