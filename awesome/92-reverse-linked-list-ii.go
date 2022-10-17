package awesome

import "fmt"

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head.Next == nil || left == right {
		return head
	}

	dummyHead := &ListNode{Val: 0, Next: head}

	indexNode := dummyHead
	var leftNode, rightNode *ListNode
	var preLeftNode, afterRightNode *ListNode
	for left >= 0 || right >= 0 {

		if left == 1 {
			preLeftNode = indexNode
			leftNode = indexNode.Next
		}
		if right == 0 {
			rightNode = indexNode
			afterRightNode = indexNode.Next
		}

		fmt.Printf("left: %v, right: %v \n", left, right)
		indexNode = indexNode.Next
		left --
		right --
	}

	 var reverseList func(head *ListNode) (*ListNode, *ListNode)
	 reverseList = func(head *ListNode) (*ListNode, *ListNode) {

		if head.Next == nil {
			return head, head
		}
		temp := head.Next
		newHead, newTail := reverseList(temp)

		newTail.Next = head
		head.Next = nil
		return newHead, head
	}

	rightNode.Next = nil
	newHead, newTail := reverseList(leftNode)

	if preLeftNode != nil {
		preLeftNode.Next = newHead
	}

	newTail.Next = afterRightNode
	return dummyHead.Next
}







