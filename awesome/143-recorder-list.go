package awesome

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	// 转成数组

	nodes := make([]*ListNode, 0, 0)

	node := head

	for ;; {
		if node == nil {
			break
		}
		nodes = append(nodes, node)
		node = node.Next
	}

	fmt.Printf("nodes: %v \n", nodes)
	n := len(nodes)

	dumpHead := &ListNode{}

	tnode := dumpHead
	for i, j := 0, n - 1 ; i <= j ; i, j = i+1, j-1 {

		fmt.Printf("i: %v, j: %v \n", i, j)

		tnode.Next = nodes[i]
		tnode = tnode.Next

		if j != i {
			tnode.Next = nodes[j]
			tnode = tnode.Next
		}

		if i + 1 > j - 1 {
			tnode.Next = nil
		}
	}

	(*head) = *dumpHead.Next

	fmt.Printf("return ")
	return

}