package awesome

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	tHead := head
	tEnd := head
	step := 0

	for ; ; {
		// 说明到最后了
		if tEnd.Next == nil {

			if step < n {
				return head.Next
			}
			if n == 1 && tHead == tEnd && tEnd.Next == nil {
				return nil
			}
			tempNode := tHead.Next
			if tempNode == nil {
				break
			}
			tHead.Next = tempNode.Next
			break
		}
		// 先走几步
		if step < n {
			tEnd = tEnd.Next
			step++
		} else {
			tEnd = tEnd.Next
			tHead = tHead.Next
		}

	}
	return head
}
