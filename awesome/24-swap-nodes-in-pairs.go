package awesome

func swapPairs(head *ListNode) *ListNode {

	var swapPairs func(head *ListNode) *ListNode
	swapPairs = func(head *ListNode) *ListNode {

		if head == nil {
			return nil
		}
		if head.Next == nil {
			return head
		}
		firstNode := head
		secondNode := firstNode.Next

		nextHead := secondNode.Next

		secondNode.Next = firstNode
		// 递归
		firstNode.Next = swapPairs(nextHead)

		return secondNode
	}
	return swapPairs(head)
}

// 迭代
func swapPairs2(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
