package awesome

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

func deleteDuplicates(head *ListNode) *ListNode {


	if head == nil {
		return nil
	}

	dummyNode := &ListNode{
		Val: -101,
		Next: head,
	}

	curNode := dummyNode

	for curNode.Next != nil && curNode.Next.Next != nil {

		// 节点值相同
		if curNode.Next.Val == curNode.Next.Next.Val {
			x := curNode.Next.Val
			// 遍历到不同为止
			for curNode.Next != nil && curNode.Next.Val == x {
				curNode.Next = curNode.Next.Next
			}


		} else {
			curNode = curNode.Next
		}
	}

	return dummyNode.Next
}