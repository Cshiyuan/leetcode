package awesome

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {


	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	// 双节点二分排序？
	midNode := head
	tailNode := head

	// 快慢指针二分
	for ;; {
		// 到达尾部
		if tailNode.Next == nil {
			break
		}
		tailNode = tailNode.Next
		// 到达尾部
		if tailNode.Next == nil {
			break
		}
		tailNode = tailNode.Next
		midNode = midNode.Next
	}

	// 中间切断
	midHead := midNode.Next
	midNode.Next = nil
	// 分归排序
	node1 := sortList(head)
	node2 := sortList(midHead)

	// 定下来新链表的头
	var newHead *ListNode
	if node1.Val > node2.Val {
		newHead = node2
		node2 = node2.Next
	} else {
		newHead = node1
		node1 = node1.Next
	}
	node := newHead
	// 遍历
	for ;; {

		if node1 == nil {
			node.Next = node2
			break
		}
		if node2 == nil {
			node.Next = node1
			break
		}
		if node1.Val > node2.Val {
			node.Next = node2
			node2 = node2.Next
		} else {
			node.Next = node1
			node1 = node1.Next
		}
		node = node.Next
	}

	return newHead
}