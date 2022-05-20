package awesome

import "math"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var root *ListNode
	var curRoot *ListNode

	curList1 := list1
	curList2 := list2

	for ; ; {
		if curList1 == nil && curList2 == nil {
			break
		}
		list1var := math.MaxInt64
		list2var := math.MaxInt64

		if curList1 != nil {
			list1var = curList1.Val
		}
		if curList2 != nil {
			list2var = curList2.Val
		}

		var tempNode *ListNode
		if list1var < list2var {
			tempNode = curList1
			curList1 = curList1.Next
		} else {
			tempNode = curList2
			curList2 = curList2.Next
		}

		// 暂时认为tempNode不可能唯恐
		if root == nil {
			root = tempNode
			curRoot = tempNode
			continue
		}
		// 拼接
		curRoot.Next = tempNode
		curRoot = curRoot.Next


	}

	return root
}
