package awesome

import "fmt"

type groupListNode struct {
	head *ListNode
	tail *ListNode
}


func reverseKGroup(head *ListNode, k int) *ListNode {


	if k == 0 || head == nil || head.Next == nil || k == 1 {
		return head
	}

	indexHead := head

	count := 0

	groupListNodes := make([]*groupListNode, 0 ,0)


	var tempHead, tempTail *ListNode

	for indexHead != nil {
		// 头
		if count % k == 0{
			tempHead = indexHead
		}
        // 下一次就是分组尾巴了，提前存储
		if (count  + 1)% k == 0 {
			tempTail = indexHead
			groupListNodes = append(groupListNodes,
				&groupListNode{
					head: tempHead,
					tail: tempTail,
				})
			// 说明分组尾巴，跟链表尾巴一致
			if indexHead.Next == nil {
				tempHead = nil
			}

		}
		count ++
		indexHead = indexHead.Next

	}
	fmt.Println(groupListNodes)

	// 将一个一个段拆开, 并反转
	for _, g := range groupListNodes {
		g.tail.Next = nil

		listNode := g.head
		g.tail = g.head
		// 反转
		g.head = reverseList(listNode)
	}


	// 将段拼接
	var tGroup *groupListNode
	for _, g := range groupListNodes {
		if tGroup == nil {
			tGroup = g
			continue
		}

		tGroup.tail.Next = g.head
		tGroup = g
	}
	tGroup.tail.Next = tempHead

	rHead := groupListNodes[0].head

	return rHead
}