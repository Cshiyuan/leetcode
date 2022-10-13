package awesome


// Definition for singly-linked list.

func reverseList2(head *ListNode) (*ListNode) {
	_, head = reverseList1(true, head)
	return head
}

func reverseList1(isFirst bool, head *ListNode) (*ListNode, *ListNode) {

	if head == nil || head.Next == nil {
		return head, head
	}

	reverseList, tailHead := reverseList1(false, head.Next)
	reverseList.Next = head
	if isFirst {
		head.Next = nil
	}
	return head, tailHead
}


func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	// 正序
	var one *ListNode = nil
	var result *ListNode = nil
	two := head
	for;; {
		result, one, two = reverseNode(one,  two)
		if two == nil {
			break
		}
	}
	if one != nil {
		return one
	}
	return result
}


func reverseNode(preHead, head *ListNode) (*ListNode, *ListNode, *ListNode) {
	// 之前的尾巴
	nextHead := head.Next
	var before *ListNode
	if nextHead != nil {
		// 之前的尾巴
		before = nextHead.Next
		// 指向
		nextHead.Next = head
	}
	head.Next = preHead
	return head, nextHead, before
}