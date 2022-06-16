package awesome

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	// 快慢指针
	slow := head
	fast := head

	for ;; {

		slow = slow.Next
		// 走两步
		fast = fast.Next
		// 没有环
		if fast == nil {
			return nil
		}
		fast = fast.Next
		// 没有环
		if fast == nil {
			return nil
		}
		fmt.Printf("slow: %v fast: %v \n", fast, slow)

		// 相遇
		if slow == fast {
			fmt.Printf("slow crash fast: %v \n", slow)
			break
		}
	}

	// 回起点
	slow = head

	for ;; {

		if slow == fast {
			return slow
		}
		// 同速率
		slow = slow.Next
		fast = fast.Next

	}

	return nil

}