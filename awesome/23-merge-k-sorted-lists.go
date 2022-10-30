package awesome

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {

	dumpNode := &ListNode{}
	curNode := dumpNode

	node := findMinElement(lists)
	for node != nil {

		curNode.Next = node
		curNode = node
		node = findMinElement(lists)
	}

	return dumpNode.Next

}
func findMinElement(nodes []*ListNode) *ListNode {

	min := math.MaxInt32
	var minNode *ListNode
	minIndex := -1

	for i := range nodes {
		if nodes[i] != nil && nodes[i].Val < min {
			minNode = nodes[i]
			min = nodes[i].Val
			minIndex = i
		}
	}

	if minIndex == -1 {
		return nil
	}

	nodes[minIndex] = nodes[minIndex].Next
	return minNode
}