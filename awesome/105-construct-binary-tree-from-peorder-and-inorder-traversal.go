package awesome

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	//pos := make(map[int]int)
	var rootPos, n int
	for rootPos, n = range inorder {
		if preorder[0] == n {
			break
		}
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	//rootPos := pos[root.Val]
	//rightPosInPreOrder := len(preorder) - 1
	//for i, v := range preorder[1:] {
	//	// 获得中序中的位置
	//	inOrderPos := pos[v]
	//	// 存储分割点的位置
	//	if inOrderPos > rootPos {
	//		rightPosInPreOrder = i
	//		break
	//	}
	//}

	root.Left = buildTree(preorder[1:len(inorder[:rootPos])+1], inorder[:rootPos])
	root.Right = buildTree(preorder[len(inorder[:rootPos])+1:], inorder[rootPos+1:])

	return root
}
