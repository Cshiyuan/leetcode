package awesome

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	var nextNodes []*TreeNode
	nextNodes = append(nextNodes, root)
	var valss [][]int
	for ; ; {
		var vals []int
		vals, nextNodes = ran_level_order(nextNodes)
		valss = append(valss, vals)
		if len(nextNodes) == 0 {
			break
		}
	}
	return valss
}

func levelOrderStd(root *TreeNode) [][]int {

	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func ran_level_order(nodes []*TreeNode) ([]int, []*TreeNode) {
	var nextNodes []*TreeNode
	var vals []int
	for _, node := range nodes {
		if node == nil {
			continue
		}
		vals = append(vals, node.Val)
		if node.Left != nil {
			nextNodes = append(nextNodes, node.Left)
		}
		if node.Right != nil {
			nextNodes = append(nextNodes, node.Right)
		}
	}
	return vals, nextNodes
}

func level_order(root *TreeNode) ([]int, [][]int) {
	if root == nil {
		return []int{}, [][]int{}
	}
	l, leftOrder := level_order(root.Left)
	r, rightOrder := level_order(root.Right)
	curOrder := append(l, r...)
	afterOrder := leftOrder
	if len(rightOrder) != 0 {
		afterOrder = append(leftOrder, rightOrder...)
	}
	var order [][]int
	if len(curOrder) != 0 {
		order = [][]int{curOrder}
	}
	if len(afterOrder) != 0 {
		order = append([][]int{curOrder}, afterOrder...)
	}

	return []int{root.Val}, order
}
